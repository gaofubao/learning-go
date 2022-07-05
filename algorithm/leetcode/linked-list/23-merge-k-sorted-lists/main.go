package main

import "container/heap"

/*
23. Merge k Sorted Lists
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
Merge all the linked-lists into one sorted linked-list and return it
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for _, list := range lists {
		if list != nil {
			heap.Push(minHeap, list)
		}
	}

	for minHeap.Len() > 0 {
		list := heap.Pop(minHeap).(*ListNode)
		curr.Next = list
		if list.Next != nil {
			heap.Push(minHeap, list.Next)
		}
		curr = curr.Next
	}

	return dummy.Next
}

// MinHeap 最小堆
type MinHeap []*ListNode

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
