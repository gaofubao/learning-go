package main

import (
	"fmt"
	"k8s.io/client-go/discovery"
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
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err)
	}

	// 发起查询请求
	groups, resourceLists, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		panic(err)
	}

	fmt.Println("group:")
	for _, group := range groups {
		fmt.Printf("\t%s\n", group.Name)
	}

	fmt.Println("resource:")
	for _, resourceList := range resourceLists {
		for _, resource := range resourceList.APIResources {
			fmt.Printf("\t%s\n", resource.Name)
		}
	}
}
