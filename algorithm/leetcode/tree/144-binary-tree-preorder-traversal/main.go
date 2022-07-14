package main

/*
144. Binary Tree Preorder Traversal
Given the root of a binary tree, return the preorder traversal of its nodes' values.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：遍历二叉树
func preorderTraversal(root *TreeNode) []int {
	var (
		res      []int
		traverse func(node *TreeNode)
	)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)
	return res
}

// 方法二：分解问题
func preorderTraversal2(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	res = append(res, root.Val)
	res = append(res, preorderTraversal2(root.Left)...)
	res = append(res, preorderTraversal2(root.Right)...)

	return res
}
