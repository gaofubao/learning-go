package main

/*
78. Subsets
Given an integer array nums of unique elements, return all possible subsets (the power set).
The solution set must not contain duplicate subsets. Return the solution in any order.
*/

func subsets(nums []int) [][]int {
	var (
		res       [][]int
		track     []int
		backtrack func(start int)
	)

	backtrack = func(start int) {
		// 前序位置
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			// 做选择
			track = append(track, nums[i])
			// 递归，进入下一层决策数，start 参数控制数的遍历，避免产生重复的子集
			backtrack(i + 1)
			// 撤销选择
			track = track[:len(track)-1]
		}
	}

	backtrack(0)
	return res
}
