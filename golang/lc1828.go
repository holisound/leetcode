package main

/*
https://leetcode.cn/problems/queries-on-number-of-points-inside-a-circle/
1828. 统计一个圆中点的数目
*/
func countPoints(points [][]int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i, q := range queries {
		for _, p := range points {
			if dist(q[0], p[0], q[1], p[1]) <= q[2]*q[2] {
				ans[i]++
			}
		}
	}
	return ans
}

func dist(a, b, c, d int) int {
	return (a-b)*(a-b) + (c-d)*(c-d)
}
