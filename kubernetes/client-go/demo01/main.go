package main

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 加载配置文件
	config, err := clientcmd.BuildConfigFromFlags("", "config.yaml")
	if err != nil {
		panic(err)
	}

	// 初始化客户端
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 实例化informer
	sharedInformers := informers.NewSharedInformerFactory(clientSet, 0)
	serviceInformer := sharedInformers.Core().V1().Services()
	ingressInformer := sharedInformers.Networking().V1().Ingresses()

	c := newController(clientSet, serviceInformer, ingressInformer)

	stopCh := make(chan struct{})
	sharedInformers.Start(stopCh)
	sharedInformers.WaitForCacheSync(stopCh)

	c.Run(stopCh)
}
