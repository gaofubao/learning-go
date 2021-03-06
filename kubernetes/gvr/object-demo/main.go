package main

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/apis/core"
	"reflect"
)

func main() {
	// 实例化一个资源对象
	pod := &core.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind: "Pod",
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{"name": "foo"},
		},
	}

	// 转换为通用资源对象
	obj := runtime.Object(pod)

	// 再转换回资源对象
	pod2, ok := obj.(*core.Pod)
	if !ok {
		panic("unexpected")
	}

	if !reflect.DeepEqual(pod, pod2) {
		panic("unexpected")
	}
}
