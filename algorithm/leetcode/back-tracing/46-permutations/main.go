package main

/*
46. Permutations
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
*/

func permute(nums []int) [][]int {
	var (
		res       [][]int
		track     []int        // 路径
		used      map[int]bool // 标记路径中的元素是否被使用过
		backtrack func()
	)

	used = make(map[int]bool)
	backtrack = func() {
		// 前置位置，只要叶子节点的值
		if len(track) == len(nums) {
			temp := make([]int, len(track))
			copy(temp, track)
			res = append(res, temp)
			return
		}

		for _, num := range nums {
			if used[num] {
				continue
			}
			// 做选择
			used[num] = true
			track = append(track, num)
			// 递归，进入下一层决策树
			backtrack()
			// 撤销选择
			used[num] = false
			track = track[:len(track)-1]
		}
	}

	backtrack()
	return res
}
