package main

/*
https://leetcode.cn/problems/last-moment-before-all-ants-fall-out-of-a-plank/
1503. 所有蚂蚁掉下来前的最后一刻
*/
func getLastMoment(n int, left []int, right []int) int {
	res := 0
	for _, x := range left {
		if x > res {
			res = x
		}
	}
	for _, x := range right {
		if n-x > res {
			res = n - x
		}
	}
	return res
}
