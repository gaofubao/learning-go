package main

import (
	"context"
	clientset "custom-resource01/pkg/generated/clientset/versioned"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "config.yaml")
	if err != nil {
		panic(err)
	}

	clientSet, err := clientset.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	fooList, err := clientSet.CrdV1().Foos("default").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		log.Println(err)
		return
	}

	for _, foo := range fooList.Items {
		fmt.Println(foo.Name)
	}
}
