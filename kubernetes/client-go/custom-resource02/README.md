# controller-tools

## 安装 controller-gen 和 type-scaffold 
```shell
git checkout v0.8.0
go install ./cmd/{controller-gen,type-scaffold}
```

## 生成 types.go
```shell
type-scaffold --kind Foo
```

## 生成 deepcopy
```shell
controller-gen object paths=./enum/apis/crd.example.com/v1/types.go
```

## 生成 crd
```shell
controller-gen crd paths=./... output:crd:dir=config/crd
```

## 编写 register.go

## 编写 main

## 参考
- https://github.com/kubernetes-sigs/controller-tools 
- https://book.kubebuilder.io/reference/controller-gen.html
- https://book.kubebuilder.io/reference/markers/crd.html
