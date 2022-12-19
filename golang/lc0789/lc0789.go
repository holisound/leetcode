package main

import (
	. "github.com/holisound/leetcode/utils"
)

/*
https://leetcode.cn/problems/escape-the-ghosts/
789. 逃脱阻碍者
*/
func escapeGhosts(ghosts [][]int, target []int) bool {
	dist := ManhattanDist(target[0], target[1], 0, 0)
	for _, g := range ghosts {
		if ManhattanDist(target[0], target[1], g[0], g[1]) <= dist {
			return false
		}
	}
	return true
}
