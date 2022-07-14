package main

import "math"

/*
104. Maximum Depth of Binary Tree
Given the root of a binary tree, return its maximum depth.
A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：遍历二叉树
func maxDepth(root *TreeNode) int {
	var (
		res      int
		depth    int
		traverse func(node *TreeNode)
	)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 前序位置
		depth++
		if node.Left == nil && node.Right == nil {
			res = int(math.Max(float64(res), float64(depth)))
		}
		traverse(node.Left)
		traverse(node.Right)
		// 后序位置
		depth--
	}

	traverse(root)
	return depth
}

// 方法二：分解问题，通过子树的最大深度推导出原树的最大深度
func maxDepth2(root *TreeNode) int {
	var res int

	if root == nil {
		return 0
	}

	leftMax := maxDepth2(root.Left)
	rightMax := maxDepth2(root.Right)
	res = int(math.Max(float64(leftMax), float64(rightMax))) + 1

	return res
}
