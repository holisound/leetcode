package main

import "sort"

/*
https://leetcode.cn/problems/sum-of-mutated-array-closest-to-target/
1300. 转变数组后最接近目标值的数组和
*/
func findBestValue(arr []int, target int) int {
	pre := 0
	n := len(arr)
	sort.Ints(arr)
	res := arr[n-1]
	for i, x := range arr {
		tot := pre + x*(n-i)
		if tot >= target {
			value := (target - pre) / (n - i)
			rem := target - pre
			if rem-value*(n-i) <= (value+1)*(n-i)-rem {
				res = value
			} else {
				res = value + 1
			}
			break
		}
		pre += x
	}
	return res
}
