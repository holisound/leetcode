package main

import "sort"

/*
https://leetcode.cn/problems/3sum/
*/
func threeSum(nums []int) [][]int {
	/*
	   这题我学到了：
	   初步了解剪枝一些用法
	*/
	sort.Ints(nums)

	res := [][]int{}
	for i, n := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			// 剪枝：对于排序数组，避免重复求解
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			s := n + nums[l] + nums[r]
			if s > 0 {
				r--
			} else if s < 0 {
				l++
			} else {
				res = append(res, []int{n, nums[l], nums[r]})
				// 为避免重复添加，l 和 r 都需要移动到不同的元素上。
				l, r = next(nums, l, r)
			}
		}
	}
	return res
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}
	return l, r
}
