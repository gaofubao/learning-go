package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 加载配置文件
	config, err := clientcmd.BuildConfigFromFlags("", "config.yaml")
	if err != nil {
		panic(err)
	}
	// 设置配置对象
	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	// 实例化客户端
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	// 发起查询请求
	res := corev1.PodList{}
	if err = restClient.
		Get().
		Namespace("kube-system").
		Resource("pods").
		VersionedParams(&metav1.ListOptions{Limit: 10}, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&res); err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range res.Items {
		fmt.Printf("%s\t%s\t%v\n", item.Namespace, item.Name, item.Status.Phase)
	}
}
