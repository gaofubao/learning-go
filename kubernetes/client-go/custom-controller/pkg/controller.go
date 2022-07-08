package pkg

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	informerCorev1 "k8s.io/client-go/informers/core/v1"
	informerNetworkingv1 "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	listerCorev1 "k8s.io/client-go/listers/core/v1"
	listerNetworkingv1 "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"reflect"
	"time"
)

const (
	WORK_NUM  = 5
	MAX_RETRY = 10
)

type controller struct {
	client        kubernetes.Interface
	serviceLister listerCorev1.ServiceLister
	ingressLister listerNetworkingv1.IngressLister
	queue         workqueue.RateLimitingInterface // 限速队列
}

func (c *controller) addService(obj interface{}) {
	c.enqueue(obj)
}

func (c *controller) updateService(oObj interface{}, nObj interface{}) {
	if reflect.DeepEqual(oObj, nObj) {
		return
	}
	c.enqueue(nObj)
}

func (c *controller) deleteIngress(obj interface{}) {
	ingress := obj.(*networkingv1.Ingress)
	service := metav1.GetControllerOf(ingress)
	if service == nil || service.Kind != "Service" {
		return
	}

	c.queue.Add(ingress.Namespace + "/" + ingress.Name)
}

func (c *controller) Run(stopCh chan struct{}) {
	for i := 0; i < WORK_NUM; i++ {
		go wait.Until(c.worker, time.Minute, stopCh)
	}
	<-stopCh
}

func (c *controller) enqueue(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
	}
	c.queue.Add(key)
}

func (c *controller) worker() {
	for c.processNextItem() {
	}
}

func (c *controller) processNextItem() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	defer c.queue.Done(item)

	key := item.(string)
	if err := c.syncService(key); err != nil {
		c.handleError(key, err)
	}
	return true
}

func (c *controller) syncService(item string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(item)
	if err != nil {
		return err
	}

	service, err := c.serviceLister.Services(namespace).Get(name)
	if errors.IsNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}

	_, ok := service.GetAnnotations()["ingress/http"]
	ingress, err := c.ingressLister.Ingresses(namespace).Get(name)
	if err != nil && !errors.IsNotFound(err) {
		return err
	}

	if !ok && ingress != nil {
		// 删除 ingress
		if err = c.client.NetworkingV1().Ingresses(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{}); err != nil {
			return err
		}

	} else if ok && errors.IsNotFound(err) {
		// 新增 ingress
		ingress = c.constructIngress(service)
		if _, err = c.client.NetworkingV1().Ingresses(namespace).Create(context.TODO(), ingress, metav1.CreateOptions{}); err != nil {
			return err
		}
	}

	return nil
}

func (c *controller) handleError(key string, err error) {
	if c.queue.NumRequeues(key) <= MAX_RETRY {
		c.queue.AddRateLimited(key)
		return
	}

	runtime.HandleError(err)
	c.queue.Forget(key)
}

/* ingress example
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ingress-myservice
spec:
  rules:
  - host: myservice.foo.org
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: myservicea
            port:
              number: 80
  ingressClassName: nginx
*/
func (c *controller) constructIngress(service *corev1.Service) *networkingv1.Ingress {
	ingress := networkingv1.Ingress{}
	// 指定属主
	ingress.ObjectMeta.OwnerReferences = []metav1.OwnerReference{
		*metav1.NewControllerRef(service, corev1.SchemeGroupVersion.WithKind("Service")),
	}
	ingress.Name = service.Name
	ingress.Namespace = service.Namespace
	ingressClass := "nginx"
	pathType := networkingv1.PathTypePrefix
	ingress.Spec = networkingv1.IngressSpec{
		IngressClassName: &ingressClass,
		Rules: []networkingv1.IngressRule{
			{
				Host: "example.com",
				IngressRuleValue: networkingv1.IngressRuleValue{
					HTTP: &networkingv1.HTTPIngressRuleValue{
						Paths: []networkingv1.HTTPIngressPath{
							{
								Path:     "/",
								PathType: &pathType,
								Backend: networkingv1.IngressBackend{
									Service: &networkingv1.IngressServiceBackend{
										Name: service.Name,
										Port: networkingv1.ServiceBackendPort{
											Number: 80,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return &ingress
}

func NewController(client kubernetes.Interface, serviceInformer informerCorev1.ServiceInformer, ingressInformer informerNetworkingv1.IngressInformer) *controller {
	c := &controller{
		client:        client,
		serviceLister: serviceInformer.Lister(),
		ingressLister: ingressInformer.Lister(),
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ingressManager"),
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
