package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 加载配置文件
	config, err := clientcmd.BuildConfigFromFlags("", "config.yaml")
	if err != nil {
		panic(err)
	}
	config.APIPath = "api"

	// 实例化客户端
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	gvr := schema.GroupVersionResource{Version: "v1", Resource: "pods"}
	unstructuredObj, err := dynamicClient.
		Resource(gvr).
		Namespace("kube-system").
		List(context.TODO(), metav1.ListOptions{Limit: 10})
	if err != nil {
		panic(err)
	}
	podList := corev1.PodList{}
	if err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredObj.UnstructuredContent(), &podList); err != nil {
		panic(err)
	}
	for _, item := range podList.Items {
		fmt.Printf("%s\t%s\t%v\n", item.Namespace, item.Name, item.Status.Phase)
	}
}
