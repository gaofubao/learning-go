package main

/*
116. Populating Next Right Pointers in Each Node
You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
Initially, all next pointers are set to NULL.
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 转换为三叉树的遍历
func connect(root *Node) *Node {
	// 遍历三叉树
	traverse(root.Left, root.Right)
	return root
}

func traverse(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}

	// 前序位置, 填充指针
	node1.Next = node2

	// 连接相同父节点的两个子节点
	traverse(node1.Left, node1.Right)
	traverse(node2.Left, node2.Right)
	// 连接跨越父节点的两个子节点
	traverse(node1.Right, node2.Left)
}
