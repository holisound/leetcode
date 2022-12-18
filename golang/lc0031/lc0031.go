package main

import "sort"

/*
https://leetcode.cn/problems/next-permutation/
*/
func nextPermutation(nums []int) {
	n := len(nums)
	index := -1
	for i := n - 2; i > -1; i-- {
		if nums[i] < nums[i+1] {
			index = i
			break
		}
	}
	if index == -1 {
		sort.Ints(nums)
	} else {
		justBigger := nums[index+1]
		index2 := index + 1
		for i := index + 2; i < n; i++ {
			if nums[i] > nums[index] && nums[i] < justBigger {
				justBigger, index2 = nums[i], i
			}

		}
		nums[index2], nums[index] = nums[index], nums[index2]
		sort.Ints(nums[index+1:])
	}

}
