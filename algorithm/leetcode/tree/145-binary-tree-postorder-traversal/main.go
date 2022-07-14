package main

/*
145. Binary Tree Postorder Traversal
Given the root of a binary tree, return the postorder traversal of its nodes' values.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：二叉树遍历
func postorderTraversal(root *TreeNode) []int {
	var (
		res      []int
		traverse func(node *TreeNode)
	)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		traverse(node.Right)
		res = append(res, node.Val)
	}

	traverse(root)
	return res
}

// 方法二：分解问题
func postorderTraversal2(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	res = append(res, postorderTraversal2(root.Left)...)
	res = append(res, postorderTraversal2(root.Right)...)
	res = append(res, root.Val)

	return res
}
