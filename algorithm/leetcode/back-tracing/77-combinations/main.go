package main

/*
77. Combinations
Given two integers n and k, return all possible combinations of k numbers out of the range [1, n].
You may return the answer in any order.
*/

func combine(n int, k int) [][]int {
	var (
		res       [][]int
		track     []int
		backtrack func(start int)
	)

	backtrack = func(start int) {
		// 前置位置，只要第k层的节点值
		if len(track) == k {
			tmp := make([]int, k)
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		for i := start; i <= n; i++ {
			track = append(track, i)
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}

	backtrack(1)
	return res
}
