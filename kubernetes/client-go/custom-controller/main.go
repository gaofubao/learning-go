package main

import (
	"github.com/gaofubao/learning-go/kubernetes/client-go/custom-controller/pkg"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 加载配置文件
	config, err := clientcmd.BuildConfigFromFlags("", "./kubernetes/client-go/custom-controller/config.yaml")
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

	// 自定义控制器
	c := pkg.NewController(clientSet, serviceInformer, ingressInformer)

	stopCh := make(chan struct{})
	// 启动 informer，监听资源变化
	sharedInformers.Start(stopCh)
	sharedInformers.WaitForCacheSync(stopCh)

	// 启动自定义控制器，对资源变化做出响应
	c.Run(stopCh)
}
