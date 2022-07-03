package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 加载配置文件
	config, err := clientcmd.BuildConfigFromFlags("", "config.yaml")
	if err != nil {
		panic(err)
	}

	// 实例化客户端
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 发起查询请求
	res, err := clientSet.
		CoreV1().
		Pods("kube-system").
		List(context.TODO(), metav1.ListOptions{Limit: 10})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range res.Items {
		fmt.Printf("%s\t%s\t%v\n", item.Namespace, item.Name, item.Status.Phase)
	}
}
