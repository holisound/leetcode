package lc1552

/*
https://leetcode.cn/problems/magnetic-force-between-two-balls/
1552. 两球之间的磁力
*/
import "sort"

// Hint: Search uses binary search to find and return the smallest index i in [0, n) at which f(i) is true
func maxDistance(position []int, m int) int {
	sort.Ints(position)
	return sort.Search((position[len(position)-1]-position[0])/(m-1), func(d int) bool {
		cnt, x0 := 1, position[0]
		for _, x := range position {
			if x >= d+1+x0 {
				cnt++
				x0 = x
			}
		}
		return cnt < m
	})
}
