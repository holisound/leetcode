package main

import (
	"math"
	"sort"
)

/*
https://leetcode.cn/problems/longest-square-streak-in-an-array/
2501. 数组中最长的方波
*/
func longestSquareStreak(nums []int) int {
	res := -1
	sort.Ints(nums)
	dp := make([]int, 100000+1)
	for _, n := range nums {
		dp[n] = 1
		m := int(math.Sqrt(float64(n)))
		if m*m != n {
			continue
		}
		if dp[m] != 0 {
			dp[n] = dp[m] + 1
			if dp[n] > res {
				res = dp[n]
			}
		}
	}

	return res
}
