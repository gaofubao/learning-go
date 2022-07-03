package main

import (
	icv1 "k8s.io/client-go/informers/core/v1"
	inv1 "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	cv1 "k8s.io/client-go/listers/core/v1"
	nv1 "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
)

type controller struct {
	client        kubernetes.Interface
	serviceLister cv1.ServiceLister
	ingressLister nv1.IngressLister
}

func (c *controller) addService(obj interface{}) {

}

func (c *controller) updateService(oObj interface{}, nObj interface{}) {

}

func (c *controller) deleteIngress(obj interface{}) {

}

func (c *controller) Run(stopCh chan struct{}) {
	<-stopCh
}

func newController(client kubernetes.Interface, serviceInformer icv1.ServiceInformer, ingressInformer inv1.IngressInformer) *controller {
	c := &controller{
		client:        client,
		serviceLister: serviceInformer.Lister(),
		ingressLister: ingressInformer.Lister(),
	}

	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addService,
		UpdateFunc: c.updateService,
	})
	ingressInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.deleteIngress,
	})

	return c
}
