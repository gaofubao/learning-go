# app-controller

## 需求
自定义资源与控制器，实现自动创建/删除 deployment/service/ingress

## 实现
### 导入 sample-controller 项目

### 按需修改

### 生成 deepcopy/clientset/listers/informers 
```shell
./hack/update-codegen.sh
```

### 编写并部署 crd/cr 
```shell
# 生成 crd 文件
controller-gen crd paths=./... output:crd:dir=artifacts/crd

# 编写 cr 文件
```

### 运行 app-controller
```shell
go run main.go --kubeconfig=/Users/xxx/.kube/config
```
观察 deployment/service/ingress 是否自动部署，
然后，删除 app，观察 deployment/service/ingress 是否自动删除。

## 参考
- https://github.com/kubernetes/sample-controller
