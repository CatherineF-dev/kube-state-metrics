/*
Copyright The Kubernetes Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	versioned "k8s.io/kube-state-metrics/pkg/customresourcemonitor/client/clientset/versioned"
	internalinterfaces "k8s.io/kube-state-metrics/pkg/customresourcemonitor/client/informers/externalversions/internalinterfaces"
	v1alpha1 "k8s.io/kube-state-metrics/pkg/customresourcemonitor/client/listers/customresourcemonitor/v1alpha1"
	customresourcemonitorv1alpha1 "k8s.io/kube-state-metrics/v2/pkg/customresourcemonitor/apis/customresourcemonitor/v1alpha1"
)

// CustomResourceMonitorInformer provides access to a shared informer and lister for
// CustomResourceMonitors.
type CustomResourceMonitorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CustomResourceMonitorLister
}

type customResourceMonitorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCustomResourceMonitorInformer constructs a new informer for CustomResourceMonitor type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCustomResourceMonitorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCustomResourceMonitorInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCustomResourceMonitorInformer constructs a new informer for CustomResourceMonitor type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCustomResourceMonitorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CustomresourceV1alpha1().CustomResourceMonitors(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CustomresourceV1alpha1().CustomResourceMonitors(namespace).Watch(context.TODO(), options)
			},
		},
		&customresourcemonitorv1alpha1.CustomResourceMonitor{},
		resyncPeriod,
		indexers,
	)
}

func (f *customResourceMonitorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCustomResourceMonitorInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *customResourceMonitorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&customresourcemonitorv1alpha1.CustomResourceMonitor{}, f.defaultInformer)
}

func (f *customResourceMonitorInformer) Lister() v1alpha1.CustomResourceMonitorLister {
	return v1alpha1.NewCustomResourceMonitorLister(f.Informer().GetIndexer())
}
