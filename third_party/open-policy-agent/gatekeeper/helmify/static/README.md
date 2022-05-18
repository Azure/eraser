# Eraser Helm Chart

## Contributing Changes

This Helm chart is autogenerated from the Eraser static manifest. The generator code lives under third_party/open-policy-agent/gatekeeper/helmify. To make modifications to this template, please edit kustomization.yaml, kustomize-for-helm.yaml and replacements.go under that directory and then run make manifests. Your changes will show up in the manifest_staging directory and will be promoted to the root charts directory the next time an Eraser release is cut.

## Get Repo Info

```console
helm repo add eraser https://azure.github.io/eraser/charts
helm repo update
```

_See [helm repo](https://helm.sh/docs/helm/helm_repo/) for command documentation._

## Install Chart

```console
# Helm install with eraser-system namespace already created
$ helm install -n eraser-system [RELEASE_NAME] eraser/eraser

# Helm install and create namespace
$ helm install -n eraser-system [RELEASE_NAME] eraser/eraser --create-namespace

```

_See [parameters](#parameters) below._

_See [helm install](https://helm.sh/docs/helm/helm_install/) for command documentation._

## Parameters

| Parameter                                     | Description                                                                                                                                                                                                                                         | Default                                                                   |
| :-------------------------------------------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------------------------------------------------------------------------ |
| image.repository                              | Image repository                                                                                                                                                                                                                                    | `ghcr.io/azure/eraser-manager`   |
| image.tag                                     | Image tag                                                                                                                                                                                                                                    | Current Chart.AppVersion |
| image.pullPolicy                              | Image pull policy                                                                                                      | `Always`   |               
| workerImage.repository                        | Image repository for worker                                                                                                                                                                                                                                    | `ghcr.io/azure/eraser`   |
| workerImage.tag                                     | Image tag for worker                                                                                                                                                                                                                                    | Current Chart.AppVersion |
| nameOverride                             | Override name if needed                                                                                                                                                                                                                | `""`                                                                      |
| securityContext                             | Security context applied on the container                                                                                                                                                                                                                | `{ allowPrivilegeEscalation: false }`                                                                      |
| resources                                     | The resource request/limits for the container image                                                                                                                                                                                                 | limits: 0.1 CPU, 30Mi, requests: 0.1 CPU, 20Mi                            |
| nodeSelector                                  | The node selector to use for pod scheduling                                                                                                                                                                                                         | `kubernetes.io/os: linux`                                                 |
| tolerations                                   | The tolerations to use for pod scheduling                                                                                                                                                                                                           | `[]`                                                                      |
| affinity                                      | The node affinity to use for pod scheduling                                                                                                                                                                                                         | `{}`                                                                      |