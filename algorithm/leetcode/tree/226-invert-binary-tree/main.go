package main

/*
226. Invert Binary Tree
Given the root of a binary tree, invert the tree, and return its root.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：遍历
func invertTree(root *TreeNode) *TreeNode {
	traverse(root)
	return root
}

func traverse(node *TreeNode) {
	if node == nil {
		return
	}

	// 前序位置, 交换左右子节点
	node.Left, node.Right = node.Right, node.Left
	// 递归左右子树
	traverse(node.Left)
	traverse(node.Right)
}

// 方法二：分解问题
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 翻转左右子树
	left := invertTree2(root.Left)
	right := invertTree2(root.Right)

	// 交换左右子节点
	root.Left, root.Right = right, left

	return root
}
