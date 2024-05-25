package main

import "sort"

// 二分查找，最大差值 最小值
func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	n := len(nums)
	return sort.Search(nums[len(nums)-1], func(mid int) bool {
		cnt := 0
		for i := 0; i+1 < n && cnt < p; {
			if nums[i+1]-nums[i] <= mid {
				cnt++
				i += 2
			} else {
				i++
			}
		}
		return cnt >= p

	})
}
