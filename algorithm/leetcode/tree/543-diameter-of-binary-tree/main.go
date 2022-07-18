package main

import "math"

/*
543. Diameter of Binary Tree
Given the root of a binary tree, return the length of the diameter of the tree.
The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
The length of a path between two nodes is represented by the number of edges between them.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var (
		res      int
		traverse func(node *TreeNode)
	)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		maxLeft := maxDepth(node.Left)
		maxRight := maxDepth(node.Right)
		res = int(math.Max(float64(res), float64(maxLeft+maxRight)))

		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)
	return res
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeft := maxDepth(root.Left)
	maxRight := maxDepth(root.Right)
	return int(math.Max(float64(maxLeft), float64(maxRight))) + 1
}
