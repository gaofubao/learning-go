package main

import (
	"context"
	v1 "custom-resource02/pkg/apis/crd.example.com/v1"
	"fmt"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "config.yaml")
	if err != nil {
		panic(err)
	}
	config.APIPath = "/apis/"
	config.GroupVersion = &v1.GroupVersion
	config.NegotiatedSerializer = v1.Codec.WithoutConversion()

	client, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	foo := v1.Foo{}
	err = client.Get().Namespace("default").Resource("foos").Name("example-foo").Do(context.TODO()).Into(&foo)
	if err != nil {
		log.Fatal(err)
	}

	foo2 := foo.DeepCopy()
	foo2.Spec.Name = "example-foo2"

	fmt.Println(foo.Spec.Name)
	fmt.Println(foo2.Spec.Name)
}
