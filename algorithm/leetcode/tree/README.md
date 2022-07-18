# 树

## 二叉树的遍历

### 递归与遍历
```shell
# 迭代遍历数组
func traverse(arr []int) {
  for i := 0; i < len(arr); i++ {
    
  }
}

# 递归遍历数组
func traverse(arr []int, i int) {
  if i == len(arr) {
    return
  }
  traverse(arr, i+1)
}

# 迭代遍历链表
func traverse(head *ListNode) {
    for p := head; p != nil; p = p.Next {
    
  }
}

# 递归遍历链表
func traverse(head *ListNode) {
  if head == nil {
    return
  }
  traverse(head.Next)
}

# 递归遍历二叉树
func traverse(root *TreeNode) {
  if root != nil {
    return
  }
  traverse(node.Left)
  traverse(node.Right)
}
```

### 前后中序遍历
- 前序遍历: [144. 二叉树的前序遍历](144-binary-tree-preorder-traversal)
- 后序遍历: [145. 二叉树的后序遍历](145-binary-tree-postorder-traversal)
- 中序遍历: [94. 二叉树的中序遍历](94-binary-tree-inorder-traversal)

前序遍历 VS 后序遍历
- 前序遍历: 自顶向下，例如与根节点相关的数据
- 后序遍历: 自底向上，例如与子树相关的数据 [543. 二叉树的直径](https://leetcode.cn/problems/diameter-of-binary-tree/)

### 递归与问题分解

### 层序遍历
[102. 二叉树的层序遍历](102-binary-tree-level-order-traversal)
