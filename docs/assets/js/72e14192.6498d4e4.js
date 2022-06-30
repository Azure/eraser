"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[239],{3905:(e,n,a)=>{a.d(n,{Zo:()=>c,kt:()=>u});var t=a(7294);function r(e,n,a){return n in e?Object.defineProperty(e,n,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[n]=a,e}function l(e,n){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var t=Object.getOwnPropertySymbols(e);n&&(t=t.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),a.push.apply(a,t)}return a}function i(e){for(var n=1;n<arguments.length;n++){var a=null!=arguments[n]?arguments[n]:{};n%2?l(Object(a),!0).forEach((function(n){r(e,n,a[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):l(Object(a)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(a,n))}))}return e}function s(e,n){if(null==e)return{};var a,t,r=function(e,n){if(null==e)return{};var a,t,r={},l=Object.keys(e);for(t=0;t<l.length;t++)a=l[t],n.indexOf(a)>=0||(r[a]=e[a]);return r}(e,n);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(t=0;t<l.length;t++)a=l[t],n.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(r[a]=e[a])}return r}var o=t.createContext({}),p=function(e){var n=t.useContext(o),a=n;return e&&(a="function"==typeof e?e(n):i(i({},n),e)),a},c=function(e){var n=p(e.components);return t.createElement(o.Provider,{value:n},e.children)},d={inlineCode:"code",wrapper:function(e){var n=e.children;return t.createElement(t.Fragment,{},n)}},m=t.forwardRef((function(e,n){var a=e.components,r=e.mdxType,l=e.originalType,o=e.parentName,c=s(e,["components","mdxType","originalType","parentName"]),m=p(a),u=r,g=m["".concat(o,".").concat(u)]||m[u]||d[u]||l;return a?t.createElement(g,i(i({ref:n},c),{},{components:a})):t.createElement(g,i({ref:n},c))}));function u(e,n){var a=arguments,r=n&&n.mdxType;if("string"==typeof e||r){var l=a.length,i=new Array(l);i[0]=m;var s={};for(var o in n)hasOwnProperty.call(n,o)&&(s[o]=n[o]);s.originalType=e,s.mdxType="string"==typeof e?e:r,i[1]=s;for(var p=2;p<l;p++)i[p]=a[p];return t.createElement.apply(null,i)}return t.createElement.apply(null,a)}m.displayName="MDXCreateElement"},4181:(e,n,a)=>{a.r(n),a.d(n,{assets:()=>o,contentTitle:()=>i,default:()=>d,frontMatter:()=>l,metadata:()=>s,toc:()=>p});var t=a(7462),r=(a(7294),a(3905));const l={title:"Quick Start",sidebar_position:3},i=void 0,s={unversionedId:"quick-start",id:"quick-start",title:"Quick Start",description:"This tutorial demonstrates the functionality of Eraser and validates that non-running images are removed succesfully.",source:"@site/docs/quick-start.md",sourceDirName:".",slug:"/quick-start",permalink:"/eraser/docs/quick-start",draft:!1,tags:[],version:"current",sidebarPosition:3,frontMatter:{title:"Quick Start",sidebar_position:3},sidebar:"tutorialSidebar",previous:{title:"Installation",permalink:"/eraser/docs/installation"},next:{title:"Development",permalink:"/eraser/docs/development"}},o={},p=[{value:"Deploy a DaemonSet",id:"deploy-a-daemonset",level:2},{value:"Excluding registries, repositories, and images",id:"excluding-registries-repositories-and-images",level:2},{value:"Automatically Cleaning Images",id:"automatically-cleaning-images",level:2},{value:"Manually Cleaning Images",id:"manually-cleaning-images",level:2}],c={toc:p};function d(e){let{components:n,...a}=e;return(0,r.kt)("wrapper",(0,t.Z)({},c,a,{components:n,mdxType:"MDXLayout"}),(0,r.kt)("p",null,"This tutorial demonstrates the functionality of Eraser and validates that non-running images are removed succesfully."),(0,r.kt)("h2",{id:"deploy-a-daemonset"},"Deploy a DaemonSet"),(0,r.kt)("p",null,"After following the ",(0,r.kt)("a",{parentName:"p",href:"https://example.com"},"install instructions"),", we'll apply a demo ",(0,r.kt)("inlineCode",{parentName:"p"},"DaemonSet"),". For illustrative purposes, a DaemonSet is applied and deleted so the non-running images remain on all nodes. The alpine image with the ",(0,r.kt)("inlineCode",{parentName:"p"},"3.7.3")," tag will be used in this example. This is an image with a known critical vulnerability."),(0,r.kt)("p",null,"First, apply the ",(0,r.kt)("inlineCode",{parentName:"p"},"DaemonSet"),":"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-shell"},"cat <<EOF | kubectl apply -f -\napiVersion: apps/v1\nkind: DaemonSet\nmetadata:\n  name: alpine\nspec:\n  selector:\n    matchLabels:\n      app: alpine\n  template:\n    metadata:\n      labels:\n        app: alpine\n    spec:\n      containers:\n      - name: alpine\n        image:docker.io/library/alpine:3.7.3\nEOF\n")),(0,r.kt)("p",null,"Next, verify that the Pods are running or completed. After the ",(0,r.kt)("inlineCode",{parentName:"p"},"alpine")," Pods complete, you may see a ",(0,r.kt)("inlineCode",{parentName:"p"},"CrashLoopBackoff")," status. This is expected behavior from the ",(0,r.kt)("inlineCode",{parentName:"p"},"alpine")," image and can be ignored for the tutorial."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-shell"},"$ kubectl get pods\nNAME          READY   STATUS      RESTARTS     AGE\nalpine-2gh9c   1/1     Running     1 (3s ago)   6s\nalpine-hljp9   0/1     Completed   1 (3s ago)   6s\n")),(0,r.kt)("p",null,"Delete the DaemonSet:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-shell"},"$ kubectl delete daemonset hello-world\n")),(0,r.kt)("p",null,"Verify that the Pods have been deleted:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-shell"},"$ kubectl get pods\nNo resources found in default namespace.\n")),(0,r.kt)("p",null,"To verify that the ",(0,r.kt)("inlineCode",{parentName:"p"},"alpine")," images are still on the nodes, exec into one of the worker nodes and list the images. If you are not using a kind cluster or Docker for your container nodes, you will need to adjust the exec command accordingly."),(0,r.kt)("p",null,"List the nodes:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-shell"},"$ kubectl get nodes\nNAME                 STATUS   ROLES           AGE   VERSION\nkind-control-plane   Ready    control-plane   45m   v1.24.0\nkind-worker          Ready    <none>          45m   v1.24.0\nkind-worker2         Ready    <none>          44m   v1.24.0\n")),(0,r.kt)("p",null,"List the images then filter for ",(0,r.kt)("inlineCode",{parentName:"p"},"alpine"),":"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-shell"},"$ docker exec kind-worker ctr -n k8s.io images list | grep alpine\ndocker.io/library/alpine:3.7.3                                                                             application/vnd.docker.distribution.manifest.list.v2+json sha256:8421d9a84432575381bfabd248f1eb56f3aa21d9d7cd2511583c68c9b7511d10 2.0 MiB   linux/386,linux/amd64,linux/arm/v6,linux/arm64/v8,linux/ppc64le,linux/s390x  io.cri-containerd.image=managed\ndocker.io/library/alpine@sha256:8421d9a84432575381bfabd248f1eb56f3aa21d9d7cd2511583c68c9b7511d10           application/vnd.docker.distribution.manifest.list.v2+json sha256:8421d9a84432575381bfabd248f1eb56f3aa21d9d7cd2511583c68c9b7511d10 2.0 MiB   linux/386,linux/amd64,linux/arm/v6,linux/arm64/v8,linux/ppc64le,linux/s390x  io.cri-containerd.image=managed\n\n")),(0,r.kt)("h2",{id:"excluding-registries-repositories-and-images"},"Excluding registries, repositories, and images"),(0,r.kt)("p",null,"Eraser can exclude registries (example, ",(0,r.kt)("inlineCode",{parentName:"p"},"docker.io/library/*"),") and also specific images with a tag (example, ",(0,r.kt)("inlineCode",{parentName:"p"},"docker.io/library/ubuntu:18.04"),") or digest (example, ",(0,r.kt)("inlineCode",{parentName:"p"},"sha256:80f31da1ac7b312ba29d65080fd..."),") from its removal process."),(0,r.kt)("p",null,"To exclude any images or registries from the removal, create a configmap named ",(0,r.kt)("inlineCode",{parentName:"p"},"excluded")," in the eraser-system namespace with a JSON file holding the excluded images."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},'$ cat > sample.json <<EOF \n{"excluded": ["docker.io/library/*", "ghcr.io/azure/test:latest"]}\nEOF\n\n$ kubectl create configmap excluded --from-file=excluded=sample.json --namespace=eraser-system\n')),(0,r.kt)("h2",{id:"automatically-cleaning-images"},"Automatically Cleaning Images"),(0,r.kt)("p",null,"After deploying Eraser, it will automatically clean images in a regular interval. This interval can be set by ",(0,r.kt)("inlineCode",{parentName:"p"},"--repeat-period")," argument to ",(0,r.kt)("inlineCode",{parentName:"p"},"eraser-controller-manager"),". The default interval is 24 hours (",(0,r.kt)("inlineCode",{parentName:"p"},"24h"),'). Valid time units are "ns", "us" (or "\xb5s"), "ms", "s", "m", "h".'),(0,r.kt)("p",null,"Eraser will schedule collector pods to each node in the cluster, and the pods will collect non-running images on those nodes. Once the collectors are done, results are de-duplicated and stored in the ",(0,r.kt)("inlineCode",{parentName:"p"},"imagecollector-shared")," object in ",(0,r.kt)("inlineCode",{parentName:"p"},"ImageCollector")," CRD."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-shell"},"$ kubectl get pods -n eraser-system\nNAME                                         READY   STATUS    RESTARTS   AGE\ncollector-kind-control-plane-lv982           1/1     Running   0          14s\ncollector-kind-control-plane-ab840           1/1     Running   0          14s\ncollector-kind-control-plane-sg352           1/1     Running   0          14s\neraser-controller-manager-649c756544-bgfds   1/1     Running   0          26s\n")),(0,r.kt)("p",null,"After collector pods are finished, scanner pod will be scheduled to each node in the cluster."),(0,r.kt)("blockquote",null,(0,r.kt)("p",{parentName:"blockquote"},"If you want to remove all the images periodically, you can skip this step by removing the ",(0,r.kt)("inlineCode",{parentName:"p"},"--scanner-image")," argument. If you are deploying with Helm, use ",(0,r.kt)("inlineCode",{parentName:"p"},'--set scanner.image.repository=""')," to remove the scanner image.")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-shell"},"$ kubectl get pods -n eraser-system\nNAME                                         READY   STATUS    RESTARTS   AGE\neraser-controller-manager-649c756544-bgfds   1/1     Running   0          36s\neraser-scanner-78p49-vxb4j                   1/1     Running   0          5s\n")),(0,r.kt)("p",null,"After scanner pods are finished, Eraser will remove the non-running images from the cluster."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-shell"},"$ kubectl get pods -n eraser-system\nNAME                                         READY   STATUS      RESTARTS   AGE\neraser-controller-manager-649c756544-bgfds   1/1     Running     0          56s\neraser-kind-control-plane-lswqn              1/1     Running     0          12s\neraser-kind-worker-wfqc                      0/1     Running     0          12s\neraser-kind-worker2-gwbit                    0/1     Running     0          12s\neraser-scanner-78p49-vxb4j                   0/1     Completed   0          25s\n")),(0,r.kt)("p",null,"Eraser pods will run to completion and the non-running images will be removed."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-shell"},"$ kubectl get pods -n eraser-system\nNAME                                         READY   STATUS      RESTARTS   AGE\neraser-controller-manager-649c756544-bgfds   1/1     Running     0          61s\neraser-kind-control-plane-lswqn              0/1     Completed   0          17s\neraser-kind-worker-wfqc                      0/1     Completed   0          17s\neraser-kind-worker2-gwbit                    0/1     Completed   0          17s\neraser-scanner-78p49-vxb4j                   0/1     Completed   0          30s\n")),(0,r.kt)("h2",{id:"manually-cleaning-images"},"Manually Cleaning Images"),(0,r.kt)("p",null,"Create an ",(0,r.kt)("inlineCode",{parentName:"p"},"ImageList")," and specify the images you would like to remove. In this case, the image ",(0,r.kt)("inlineCode",{parentName:"p"},"docker.io/library/alpine:3.7.3")," will be removed."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-shell"},'cat <<EOF | kubectl apply -f -\napiVersion: eraser.sh/v1alpha1\nkind: ImageList\nmetadata:\n  name: imagelist\nspec:\n  images:\n    - docker.io/library/alpine:3.7.3   # use "*" for all non-running images\nEOF\n')),(0,r.kt)("blockquote",null,(0,r.kt)("p",{parentName:"blockquote"},(0,r.kt)("inlineCode",{parentName:"p"},"ImageList")," is a cluster-scoped resource and must be called imagelist. ",(0,r.kt)("inlineCode",{parentName:"p"},'"*"')," can be specified to remove all non-running images instead of individual images.")),(0,r.kt)("p",null,"Creating an ",(0,r.kt)("inlineCode",{parentName:"p"},"ImageList")," should trigger an ",(0,r.kt)("inlineCode",{parentName:"p"},"ImageJob")," that will deploy Eraser pods on every node to perform the removal given the list of images."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-shell"},"$ kubectl get pods -n eraser-system\neraser-system        eraser-controller-manager-55d54c4fb6-dcglq   1/1     Running   0          9m8s\neraser-system        eraser-kind-control-plane                    1/1     Running   0          11s\neraser-system        eraser-kind-worker                           1/1     Running   0          11s\neraser-system        eraser-kind-worker2                          1/1     Running   0          11s\n")),(0,r.kt)("p",null,"Pods will run to completion and the images will be removed."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-shell"},"$ kubectl get pods -n eraser-system\neraser-system        eraser-controller-manager-6d6d5594d4-phl2q   1/1     Running     0          4m16s\neraser-system        eraser-kind-control-plane                    0/1     Completed   0          22s\neraser-system        eraser-kind-worker                           0/1     Completed   0          22s\neraser-system        eraser-kind-worker2                          0/1     Completed   0          22s\n")),(0,r.kt)("p",null,"The ",(0,r.kt)("inlineCode",{parentName:"p"},"ImageList")," custom resource status field will contain the status of the last job. The success and failure counts indicate the number of nodes the Eraser agent was run on."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-shell"},"$ kubectl describe ImageList imagelist\n...\nStatus:\n  Failed:     0\n  Success:    3\n  Timestamp:  2022-02-25T23:41:55Z\n...\n")),(0,r.kt)("p",null,"Verify the unused images are removed."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-shell"},"$ docker exec kind-worker ctr -n k8s.io images list | grep alpine\n")),(0,r.kt)("p",null,"If the image has been successfully removed, there will be no output."))}d.isMDXComponent=!0}}]);