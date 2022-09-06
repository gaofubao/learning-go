# code-generator

## 定义 doc.go/types.go/register.go 文件

## 生成代码
```shell
 ~/go/src/k8s.io/code-generator/generate-groups.sh all custom-resource01/enum/generated custom-resource01/enum/apis crd.example.com:v1 --go-header-file=/Users/gaofubao/go/src/k8s.io/code-generator/hack/boilerplate.go.txt --output-base ../
```

## 编写 main 文件

## 参考
- https://github.com/kubernetes/code-generator
- https://github.com/kubernetes/sample-controller
