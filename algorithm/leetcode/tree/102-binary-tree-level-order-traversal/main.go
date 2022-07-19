package main

/*
102. Binary Tree Level Order Traversal
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	levelQueue := []*TreeNode{root}
	for len(levelQueue) != 0 {
		var (
			tmpValue []int
			tmpQueue []*TreeNode
		)

		for _, node := range levelQueue {
			tmpValue = append(tmpValue, node.Val)

			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}
			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}
		}
		res = append(res, tmpValue)
		levelQueue = tmpQueue
	}

	return res
}
