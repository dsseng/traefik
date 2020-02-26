/*
The MIT License (MIT)

Copyright (c) 2016-2020 Containous SAS

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	versioned "github.com/containous/traefik/v2/pkg/provider/kubernetes/crd/generated/clientset/versioned"
	internalinterfaces "github.com/containous/traefik/v2/pkg/provider/kubernetes/crd/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/containous/traefik/v2/pkg/provider/kubernetes/crd/generated/listers/traefik/v1alpha1"
	traefikv1alpha1 "github.com/containous/traefik/v2/pkg/provider/kubernetes/crd/traefik/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MiddlewareInformer provides access to a shared informer and lister for
// Middlewares.
type MiddlewareInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MiddlewareLister
}

type middlewareInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMiddlewareInformer constructs a new informer for Middleware type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMiddlewareInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMiddlewareInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMiddlewareInformer constructs a new informer for Middleware type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMiddlewareInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TraefikV1alpha1().Middlewares(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TraefikV1alpha1().Middlewares(namespace).Watch(options)
			},
		},
		&traefikv1alpha1.Middleware{},
		resyncPeriod,
		indexers,
	)
}

func (f *middlewareInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMiddlewareInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *middlewareInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&traefikv1alpha1.Middleware{}, f.defaultInformer)
}

func (f *middlewareInformer) Lister() v1alpha1.MiddlewareLister {
	return v1alpha1.NewMiddlewareLister(f.Informer().GetIndexer())
}
