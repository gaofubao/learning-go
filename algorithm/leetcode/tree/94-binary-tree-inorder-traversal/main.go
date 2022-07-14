package main

/*
94. Binary Tree Inorder Traversal
Given the root of a binary tree, return the inorder traversal of its nodes' values.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：遍历二叉树
func inorderTraversal(root *TreeNode) []int {
	var (
		res      []int
		traverse func(node *TreeNode)
	)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		res = append(res, node.Val)
		traverse(node.Right)
	}

	traverse(root)
	return res
}

// 方法二：分解问题
func inorderTraversal2(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	res = append(res, inorderTraversal2(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal2(root.Right)...)

	return res
}
