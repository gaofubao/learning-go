package main

/*
86. Partition List
Given the head of a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
You should preserve the original relative order of the nodes in each of the two partitions.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	// 定义两个哨兵节点分别指向小于x的节点和大于等于x的节点
	dummy1, dummy2 := &ListNode{}, &ListNode{}
	curr1, curr2 := dummy1, dummy2

	// 遍历原链表，拆分为两个链表
	for head != nil {
		if head.Val < x {
			curr1.Next = head
			curr1 = curr1.Next
		} else {
			curr2.Next = head
			curr2 = curr2.Next
		}
		// 切断原链表的 Next 指针
		head.Next, head = nil, head.Next
	}

	// 连接两个链表
	curr1.Next = dummy2.Next
	return dummy1.Next
}
