package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	machinerytypes "k8s.io/apimachinery/pkg/types"

	eraserv1alpha1 "github.com/Azure/eraser/api/v1alpha1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/aquasecurity/fanal/analyzer/config"
	"github.com/aquasecurity/fanal/applier"
	"github.com/aquasecurity/fanal/artifact"
	artifactImage "github.com/aquasecurity/fanal/artifact/image"
	"github.com/aquasecurity/fanal/cache"
	fanalImage "github.com/aquasecurity/fanal/image"
	fanalTypes "github.com/aquasecurity/fanal/types"

	"github.com/aquasecurity/trivy-db/pkg/db"

	dlDb "github.com/aquasecurity/trivy/pkg/db"
	"github.com/aquasecurity/trivy/pkg/detector/ospkg"
	pkgResult "github.com/aquasecurity/trivy/pkg/result"
	"github.com/aquasecurity/trivy/pkg/scanner"
	"github.com/aquasecurity/trivy/pkg/scanner/local"
	trivyTypes "github.com/aquasecurity/trivy/pkg/types"
)

const (
	generalErr = 1

	severityCritical = "CRITICAL"
	severityHigh     = "HIGH"
	severityMedium   = "MEDIUM"
	severityLow      = "LOW"
	severityUnknown  = "UNKNOWN"

	vulnTypeOs      = "os"
	vulnTypeLibrary = "library"

	securityCheckVuln   = "vuln"
	securityCheckConfig = "config"
	securityCheckSecret = "secret"

	apiPath         = "apis/eraser.sh/v1alpha1"
	resourceName    = "imagecollectors"
	subResourceName = "status"
)

var (
	collectorCRName = flag.String("collector-cr-name", "collector-cr", "name of the collector cr to read from and write to")
	cacheDir        = flag.String("cache-dir", "/var/lib/trivy", "path to the cache dir")
	severity        = flag.String("severity", "CRITICAL", "list of severity levels to report")
	ignoreUnfixed   = flag.Bool("ignore-unfixed", true, "report only fixed vulnerabilities")
	vulnTypes       = flag.String("vuln-type", "os,library", "comma separated list of vulnerability types")
	securityChecks  = flag.String("security-checks", "vuln,secret", "comma-separated list of what security issues to detect")

	// Will be modified by parseCommaSeparatedOptions() to reflect the `severity` CLI flag
	// These are the only recognized severities and the keys of this map should never be modified.
	severityMap = map[string]bool{
		severityCritical: false,
		severityHigh:     false,
		severityMedium:   false,
		severityLow:      false,
		severityUnknown:  false,
	}

	// Will be modified by parseCommaSeparatedOptions() to reflect the `security-checks` CLI flag
	// These are the only recognized security checks and the keys of this map should never be modified.
	securityCheckMap = map[string]bool{
		securityCheckVuln:   false,
		securityCheckSecret: false,
		securityCheckConfig: false,
	}

	// Will be modified by parseCommaSeparatedOptions()  to reflect the `vuln-type` CLI flag
	// These are the only recognized vulnerability types and the keys of this map should never be modified.
	vulnTypeMap = map[string]bool{
		vulnTypeOs:      false,
		vulnTypeLibrary: false,
	}
)

type (
	scannerSetup struct {
		fscache       cache.FSCache
		localScanner  local.Scanner
		scanOptions   trivyTypes.ScanOptions
		dockerOptions fanalTypes.DockerOption
	}

	patch struct {
		Status eraserv1alpha1.ImageCollectorStatus `json:"status"`
	}

	optionSet struct {
		input string
		m     map[string]bool
	}

	statusUpdate struct {
		apiPath         string
		ctx             context.Context
		clientset       *kubernetes.Clientset
		collectorCRName string
		resourceName    string
		subResourceName string
		images          []eraserv1alpha1.Image
	}

	imageErr struct {
		eraserv1alpha1.Image
		err error
	}
)

func main() {
	flag.Parse()
	ctx := context.Background()

	allCommaSeparatedOptions := []optionSet{
		{
			input: *severity,
			m:     severityMap,
		},
		{
			input: *vulnTypes,
			m:     vulnTypeMap,
		},
		{
			input: *securityChecks,
			m:     securityCheckMap,
		},
	}

	for _, oSet := range allCommaSeparatedOptions {
		// note: this function has side effects and will modify the map supplied as the first argument
		err := parseCommaSeparatedOptions(oSet.m, oSet.input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(generalErr)
		}
	}

	cfg, err := rest.InClusterConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(generalErr)
	}

	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(generalErr)
	}

	result := eraserv1alpha1.ImageCollector{}

	err = clientset.RESTClient().Get().
		AbsPath(apiPath).
		Resource(resourceName).
		Name(*collectorCRName).
		Do(context.Background()).
		Into(&result)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(generalErr)
	}

	// deduplicate, which may not be necessary
	scanList := make(map[string]string)
	for _, img := range result.Spec.Images {
		scanList[img.Digest] = img.Name
	}

	err = downloadAndInitDB(*cacheDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(generalErr)
	}

	vulnTypeList := trueMapKeys(vulnTypeMap)
	securityCheckList := trueMapKeys(securityCheckMap)

	scanConfig, err := setupScanner(*cacheDir, vulnTypeList, securityCheckList)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(generalErr)
	}

	resultClient := initializeResultClient()
	resultClientMutex := sync.Mutex{}
	imgChan := make(chan imageErr)
	var wg sync.WaitGroup

	for k := range result.Spec.Images {
		img := result.Spec.Images[k]
		imageName := img.Name
		wg.Add(1)

		go func(imageRef string, img eraserv1alpha1.Image, resultClientMutex *sync.Mutex) {
			if imageRef == "" {
				imgChan <- imageErr{Image: img, err: fmt.Errorf("no name")}
				wg.Done()
				return
			}

			fmt.Printf("scanning: %s\n", imageRef)
			resultClientMutex.Lock()
			dockerImage, cleanup, err := fanalImage.NewDockerImage(ctx, imageRef, scanConfig.dockerOptions)
			resultClientMutex.Unlock()
			defer cleanup()
			if err != nil {
				imgChan <- imageErr{Image: img, err: err}
				fmt.Fprintln(os.Stderr, err)
				wg.Done()
				return
			}

			artifactToScan, err := artifactImage.NewArtifact(dockerImage, scanConfig.fscache, artifact.Option{}, config.ScannerOption{})
			if err != nil {
				imgChan <- imageErr{Image: img, err: err}
				fmt.Fprintln(os.Stderr, err)
				wg.Done()
				return
			}

			scanner := scanner.NewScanner(scanConfig.localScanner, artifactToScan)
			report, err := scanner.ScanArtifact(ctx, scanConfig.scanOptions)
			if err != nil {
				imgChan <- imageErr{Image: img, err: err}
				fmt.Fprintln(os.Stderr, err)
				wg.Done()
				return
			}

		outer:
			for i := range report.Results {
				resultClient.FillVulnerabilityInfo(report.Results[i].Vulnerabilities, report.Results[i].Type)

				for j := range report.Results[i].Vulnerabilities {
					if *ignoreUnfixed && report.Results[i].Vulnerabilities[j].FixedVersion == "" {
						continue
					}

					if report.Results[i].Vulnerabilities[j].Severity == "" {
						report.Results[i].Vulnerabilities[j].Severity = severityUnknown
					}

					if severityMap[report.Results[i].Vulnerabilities[j].Severity] {
						imgChan <- imageErr{Image: img, err: nil}
						break outer
					}
				}
			}

			wg.Done()
		}(imageName, img, &resultClientMutex)
	}

	go func() {
		wg.Wait()
		close(imgChan)
	}()

	vulnerableImages := make([]eraserv1alpha1.Image, 0, len(result.Spec.Images))
	for img := range imgChan {
		img.ScanSucceeded = boolPtr(true)
		if img.err != nil {
			img.ScanSucceeded = boolPtr(false)
		}
		vulnerableImages = append(vulnerableImages, img.Image)
	}

	err = updateStatus(&statusUpdate{
		apiPath:         apiPath,
		ctx:             ctx,
		clientset:       clientset,
		collectorCRName: *collectorCRName,
		resourceName:    resourceName,
		subResourceName: subResourceName,
		images:          vulnerableImages,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(generalErr)
	}

	for _, imageRef := range vulnerableImages {
		fmt.Println(imageRef)
	}
}

func boolPtr(b bool) *bool {
	return &b
}

// side effects: map `m` will be modified according to the values in `commaSeparatedList`.
func parseCommaSeparatedOptions(m map[string]bool, commaSeparatedList string) error {
	list := strings.Split(commaSeparatedList, ",")
	for _, item := range list {
		if _, ok := m[item]; !ok {
			keys := mapKeys(m)
			return fmt.Errorf("'%s' was not one of %#v", item, keys)
		}

		m[item] = true
	}

	return nil
}

func downloadAndInitDB(cacheDir string) error {
	err := downloadDB(cacheDir)
	if err != nil {
		return err
	}

	err = db.Init(cacheDir)
	if err != nil {
		return err
	}

	return nil
}

func downloadDB(cacheDir string) error {
	client := dlDb.NewClient(cacheDir, true)
	ctx := context.Background()
	needsUpdate, err := client.NeedsUpdate("dev", false)
	if err != nil {
		return err
	}

	if needsUpdate {
		if err = client.Download(ctx, cacheDir); err != nil {
			return err
		}
	}

	return nil
}

func setupScanner(cacheDir string, vulnTypes, securityChecks []string) (scannerSetup, error) {
	filesystemCache, err := cache.NewFSCache(cacheDir)
	if err != nil {
		return scannerSetup{}, err
	}

	app := applier.NewApplier(filesystemCache)
	det := ospkg.Detector{}
	dopts := fanalTypes.DockerOption{}
	scan := local.NewScanner(app, det)

	sopts := trivyTypes.ScanOptions{
		VulnType:            vulnTypes,
		SecurityChecks:      securityChecks,
		ScanRemovedPackages: false,
		ListAllPackages:     false,
	}

	return scannerSetup{
		localScanner:  scan,
		scanOptions:   sopts,
		dockerOptions: dopts,
		fscache:       filesystemCache,
	}, nil
}

func initializeResultClient() pkgResult.Client {
	config := db.Config{}
	client := pkgResult.NewClient(config)
	return client
}

func updateStatus(opts *statusUpdate) error {
	collectorPatch := patch{
		Status: eraserv1alpha1.ImageCollectorStatus{
			Result: opts.images,
		},
	}

	body, err := json.Marshal(&collectorPatch)
	if err != nil {
		return err
	}

	_, err = opts.clientset.RESTClient().Patch(machinerytypes.MergePatchType).
		AbsPath(opts.apiPath).
		Resource(opts.resourceName).
		SubResource(opts.subResourceName).
		Name(opts.collectorCRName).
		Body(body).DoRaw(opts.ctx)

	return err
}

func mapKeys(m map[string]bool) []string {
	list := []string{}
	for k := range m {
		list = append(list, k)
	}

	return list
}

func trueMapKeys(m map[string]bool) []string {
	list := []string{}
	for k := range m {
		if m[k] {
			list = append(list, k)
		}
	}

	return list
}
