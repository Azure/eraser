/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	erasershv1alpha1 "github.com/Azure/eraser/api/eraser.sh/v1alpha1"
	versioned "github.com/Azure/eraser/pkg/client/clientset/versioned"
	internalinterfaces "github.com/Azure/eraser/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/Azure/eraser/pkg/client/listers/eraser.sh/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ImageCollectorInformer provides access to a shared informer and lister for
// ImageCollectors.
type ImageCollectorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ImageCollectorLister
}

type imageCollectorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewImageCollectorInformer constructs a new informer for ImageCollector type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewImageCollectorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredImageCollectorInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredImageCollectorInformer constructs a new informer for ImageCollector type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredImageCollectorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EraserV1alpha1().ImageCollectors(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EraserV1alpha1().ImageCollectors(namespace).Watch(context.TODO(), options)
			},
		},
		&erasershv1alpha1.ImageCollector{},
		resyncPeriod,
		indexers,
	)
}

func (f *imageCollectorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredImageCollectorInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *imageCollectorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&erasershv1alpha1.ImageCollector{}, f.defaultInformer)
}

func (f *imageCollectorInformer) Lister() v1alpha1.ImageCollectorLister {
	return v1alpha1.NewImageCollectorLister(f.Informer().GetIndexer())
}
