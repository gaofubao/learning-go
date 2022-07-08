package main

import "sort"

/*
698. Partition to K Equal Sum Subsets
Given an integer array nums and an integer k, return true if it is possible to divide this array into k non-empty subsets whose sums are all equal.
*/

// 方法一
// 以数字的视角，将 n 个数字分配到 k 个桶里，每个数字都要选择进入到 k 个桶中的某一个
// 时间复杂度较高
func canPartitionKSubsets(nums []int, k int) bool {
	// 排出一些基本情况
	if k > len(nums) {
		return false
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}

	// 准备 k 个桶
	bucket := make([]int, k)
	// 每个桶的数字之和应为 target
	target := sum / k
	// 逆序排列 nums，使得尽可能多地命中剪枝
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	return backtrack(nums, 0, bucket, target)
}

// 递归穷举 nums 中的每个数字
func backtrack(nums []int, index int, bucket []int, target int) bool {
	// 结束条件
	if index == len(nums) {
		for i := 0; i < len(bucket); i++ {
			if bucket[i] != target {
				return false
			}
		}
		return true
	}

	// 穷举 nums[index] 可能装入的桶
	for i := 0; i < len(bucket); i++ {
		// 剪枝，桶装满了
		if bucket[i]+nums[index] > target {
			continue
		}

		// 做出选择
		bucket[i] += nums[index]
		// 进入下一个数字的选择
		if backtrack(nums, index+1, bucket, target) {
			return true
		}
		// 撤销选择
		bucket[i] -= nums[index]
	}

	// nums[index] 装入哪个桶都不行
	return false
}

// 方法二
// 以桶的视角，对于每个桶，都要遍历 nums 中的 n 个数字，然后选择是否将当前遍历到的数字装进自己的桶里
func canPartitionKSubsets2(nums []int, k int) bool {
	// 排出一些基本情况
	if k > len(nums) {
		return false
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}

	// 数字是否已被使用，这里使用位图
	used := 0
	// 记录当前状态是否出现过，避免重复计算
	memo := make(map[int]bool)
	// 每个桶的数字之和应为 target
	target := sum / k

	return backtrack2(nums, 0, k, 0, used, memo, target)
}

func backtrack2(nums []int, start int, k int, bucket int, used int, memo map[int]bool, target int) bool {
	// 所有桶都装满了
	if k == 0 {
		return true
	}

	// 当前桶装满了，开始装下一个桶
	if bucket == target {
		res := backtrack2(nums, 0, k-1, 0, used, memo, target)
		memo[used] = res
		return res
	}

	if res, ok := memo[used]; ok {
		return res
	}

	for i := start; i < len(nums); i++ {
		// 剪枝，判断第 i 位是否为 1
		if ((used >> i) & 1) == 1 {
			continue
		}

		// 当前桶装不下 nums[i]
		if nums[i]+bucket > target {
			continue
		}

		// 做选择
		used |= 1 << i
		bucket += nums[i]
		// 递归穷举下一个数字是否装入当前桶
		if backtrack2(nums, i+1, k, bucket, used, memo, target) {
			return true
		}
		// 撤销选择
		used ^= 1 << i
		bucket -= nums[i]
	}

	// 穷举了所有数字，都无法装满当前桶
	return false
}
