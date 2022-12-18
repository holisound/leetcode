package main

import "sort"

/*
https://leetcode.cn/problems/find-the-kth-largest-integer-in-the-array/
1985. 找出数组中的第 K 大整数
*/
func kthLargestNumber(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j] // i, j 可以互换
		if len(x) == len(y) {
			r := false
			for k := 0; k < len(x); k++ {
				if x[k] < y[k] {
					r = true
				}
				if x[k] != y[k] {
					break
				}
			}
			return r
		}
		return len(x) < len(y)
	})
	return nums[len(nums)-k]
}
