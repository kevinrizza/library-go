// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	apioperatorv1 "github.com/openshift/api/operator/v1"
	versioned "github.com/openshift/client-go/operator/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/operator/informers/externalversions/internalinterfaces"
	operatorv1 "github.com/openshift/client-go/operator/listers/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ConsoleInformer provides access to a shared informer and lister for
// Consoles.
type ConsoleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() operatorv1.ConsoleLister
}

type consoleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewConsoleInformer constructs a new informer for Console type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewConsoleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredConsoleInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredConsoleInformer constructs a new informer for Console type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredConsoleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().Consoles().List(context.Background(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().Consoles().Watch(context.Background(), options)
			},
			ListWithContextFunc: func(ctx context.Context, options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().Consoles().List(ctx, options)
			},
			WatchFuncWithContext: func(ctx context.Context, options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().Consoles().Watch(ctx, options)
			},
		},
		&apioperatorv1.Console{},
		resyncPeriod,
		indexers,
	)
}

func (f *consoleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredConsoleInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *consoleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apioperatorv1.Console{}, f.defaultInformer)
}

func (f *consoleInformer) Lister() operatorv1.ConsoleLister {
	return operatorv1.NewConsoleLister(f.Informer().GetIndexer())
}
