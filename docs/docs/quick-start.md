---
title: Quick Start
---

This tutorial demonstrates the functionality of Eraser and validates that non-running images are removed succesfully.

## Deploy a DaemonSet

After following the [install instructions](installation.md), we'll apply a demo `DaemonSet`. For illustrative purposes, a DaemonSet is applied and deleted so the non-running images remain on all nodes. The alpine image with the `3.7.3` tag will be used in this example. This is an image with a known critical vulnerability.

First, apply the `DaemonSet`:

```shell
cat <<EOF | kubectl apply -f -
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: alpine
spec:
  selector:
    matchLabels:
      app: alpine
  template:
    metadata:
      labels:
        app: alpine
    spec:
      containers:
      - name: alpine
        image:docker.io/library/alpine:3.7.3
EOF
```

Next, verify that the Pods are running or completed. After the `alpine` Pods complete, you may see a `CrashLoopBackoff` status. This is expected behavior from the `alpine` image and can be ignored for the tutorial.

```shell
$ kubectl get pods
NAME          READY   STATUS      RESTARTS     AGE
alpine-2gh9c   1/1     Running     1 (3s ago)   6s
alpine-hljp9   0/1     Completed   1 (3s ago)   6s
```

Delete the DaemonSet:

```shell
$ kubectl delete daemonset alpine
```

Verify that the Pods have been deleted:

```shell
$ kubectl get pods
No resources found in default namespace.
```

To verify that the `alpine` images are still on the nodes, exec into one of the worker nodes and list the images. If you are not using a kind cluster or Docker for your container nodes, you will need to adjust the exec command accordingly.

List the nodes:

```shell
$ kubectl get nodes
NAME                 STATUS   ROLES           AGE   VERSION
kind-control-plane   Ready    control-plane   45m   v1.24.0
kind-worker          Ready    <none>          45m   v1.24.0
kind-worker2         Ready    <none>          44m   v1.24.0
```

List the images then filter for `alpine`:

```shell
$ docker exec kind-worker ctr -n k8s.io images list | grep alpine
docker.io/library/alpine:3.7.3                                                                             application/vnd.docker.distribution.manifest.list.v2+json sha256:8421d9a84432575381bfabd248f1eb56f3aa21d9d7cd2511583c68c9b7511d10 2.0 MiB   linux/386,linux/amd64,linux/arm/v6,linux/arm64/v8,linux/ppc64le,linux/s390x  io.cri-containerd.image=managed
docker.io/library/alpine@sha256:8421d9a84432575381bfabd248f1eb56f3aa21d9d7cd2511583c68c9b7511d10           application/vnd.docker.distribution.manifest.list.v2+json sha256:8421d9a84432575381bfabd248f1eb56f3aa21d9d7cd2511583c68c9b7511d10 2.0 MiB   linux/386,linux/amd64,linux/arm/v6,linux/arm64/v8,linux/ppc64le,linux/s390x  io.cri-containerd.image=managed

```

## Automatically Cleaning Images

After deploying Eraser, it will automatically clean images in a regular interval. This interval can be set by `--repeat-period` argument to `eraser-controller-manager`. The default interval is 24 hours (`24h`). Valid time units are "ns", "us" (or "??s"), "ms", "s", "m", "h".

Eraser will schedule collector pods to each node in the cluster, and the pods will collect non-running images on those nodes. Once the collectors are done, results are de-duplicated and stored in the `imagecollector-shared` object in `ImageCollector` CRD.

```shell
$ kubectl get pods -n eraser-system
NAME                                         READY   STATUS    RESTARTS   AGE
collector-kind-control-plane-lv982           1/1     Running   0          14s
collector-kind-control-plane-ab840           1/1     Running   0          14s
collector-kind-control-plane-sg352           1/1     Running   0          14s
eraser-controller-manager-649c756544-bgfds   1/1     Running   0          26s
```

After collector pods are finished, scanner pod will be scheduled to each node in the cluster.

> If you want to remove all the images periodically, you can skip this step by removing the `--scanner-image` argument. If you are deploying with Helm, use `--set scanner.image.repository=""` to remove the scanner image.

```shell
$ kubectl get pods -n eraser-system
NAME                                         READY   STATUS    RESTARTS   AGE
eraser-controller-manager-649c756544-bgfds   1/1     Running   0          36s
eraser-scanner-78p49-vxb4j                   1/1     Running   0          5s
```

After scanner pods are finished, Eraser will remove the non-running images from the cluster.

```shell
$ kubectl get pods -n eraser-system
NAME                                         READY   STATUS      RESTARTS   AGE
eraser-controller-manager-649c756544-bgfds   1/1     Running     0          56s
eraser-kind-control-plane-lswqn              1/1     Running     0          12s
eraser-kind-worker-wfqc                      0/1     Running     0          12s
eraser-kind-worker2-gwbit                    0/1     Running     0          12s
eraser-scanner-78p49-vxb4j                   0/1     Completed   0          25s
```

Eraser pods will run to completion and the non-running images will be removed.

```shell
$ kubectl get pods -n eraser-system
NAME                                         READY   STATUS      RESTARTS   AGE
eraser-controller-manager-649c756544-bgfds   1/1     Running     0          61s
eraser-kind-control-plane-lswqn              0/1     Completed   0          17s
eraser-kind-worker-wfqc                      0/1     Completed   0          17s
eraser-kind-worker2-gwbit                    0/1     Completed   0          17s
eraser-scanner-78p49-vxb4j                   0/1     Completed   0          30s
```
