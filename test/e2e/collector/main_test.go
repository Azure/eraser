//go:build collector
// +build collector

package collector

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	eraserv1alpha1 "github.com/Azure/eraser/api/v1alpha1"
	"github.com/Azure/eraser/test/e2e/util"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/klog/v2"
	"sigs.k8s.io/e2e-framework/klient/wait"
	"sigs.k8s.io/e2e-framework/klient/wait/conditions"
	"sigs.k8s.io/e2e-framework/pkg/env"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/envfuncs"
)

var (
	providerResourceDirectory = "manifest_staging/charts"
	providerResource          = "eraser.yaml"
	eraserNamespace           = "eraser-system"
	testenv                   env.Environment
	image                     = os.Getenv("IMAGE")
	managerImage              = os.Getenv("MANAGER_IMAGE")
	collectorImage            = os.Getenv("COLLECTOR_IMAGE")
	scannerImage              = os.Getenv("SCANNER_IMAGE")
	vulnerableImage           = os.Getenv("VULNERABLE_IMAGE")
	nodeVersion               = os.Getenv("NODE_VERSION")
)

func TestMain(m *testing.M) {
	utilruntime.Must(eraserv1alpha1.AddToScheme(scheme.Scheme))

	testenv = env.NewWithConfig(envconf.New())
	// Create KinD Cluster
	namespace := envconf.RandomName("eraser-ns", 16)
	testenv.Setup(
		envfuncs.CreateKindClusterWithConfig(util.KindClusterName, nodeVersion, "../kind-config.yaml"),
		envfuncs.CreateNamespace(namespace),
		envfuncs.LoadDockerImageToCluster(util.KindClusterName, managerImage),
		envfuncs.LoadDockerImageToCluster(util.KindClusterName, image),
		envfuncs.LoadDockerImageToCluster(util.KindClusterName, scannerImage),
		envfuncs.LoadDockerImageToCluster(util.KindClusterName, collectorImage),
		envfuncs.LoadDockerImageToCluster(util.KindClusterName, vulnerableImage),
		deployEraserManifest(eraserNamespace),
	).Finish(
		envfuncs.DeleteNamespace(namespace),
	)
	os.Exit(testenv.Run(m))
}

func deployEraserManifest(namespace string) env.Func {
	return func(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
		wd, err := os.Getwd()
		if err != nil {
			return ctx, err
		}

		providerResourceAbsolutePath, err := filepath.Abs(filepath.Join(wd, "/../../../", providerResourceDirectory, "eraser"))
		if err != nil {
			return ctx, err
		}
		// start deployment
		if err := util.HelmInstall(cfg.KubeconfigFile(), namespace, []string{providerResourceAbsolutePath}); err != nil {
			return ctx, err
		}

		client, err := cfg.NewClient()
		if err != nil {
			klog.ErrorS(err, "Failed to create new Client")
			return ctx, err
		}

		// wait for the deployment to finish becoming available
		eraserManagerDep := appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{Name: "eraser-controller-manager", Namespace: namespace},
		}

		if err = wait.For(conditions.New(client.Resources()).DeploymentConditionMatch(&eraserManagerDep, appsv1.DeploymentAvailable, corev1.ConditionTrue),
			wait.WithTimeout(time.Minute*1)); err != nil {
			klog.ErrorS(err, "failed to deploy eraser manager")

			return ctx, err
		}

		return ctx, nil
	}
}