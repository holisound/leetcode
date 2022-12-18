package main

import (
	. "github.com/holisound/leetcode/utils"
)

/*
https://leetcode.cn/problems/path-with-maximum-gold/
1219. 黄金矿工
*/
var (
	res, m, n int
	dx        []int = []int{0, 0, 1, -1}
	dy        []int = []int{1, -1, 0, 0}
)

func getMaximumGold(grid [][]int) int {
	res = 0
	n, m = len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] != 0 {
				res = Max(res, dfs(grid, i, j))
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int) int {
	if grid[i][j] == 0 {
		return 0
	}
	gold := grid[i][j]
	grid[i][j] = 0
	tot := 0
	for k := 0; k < 4; k++ {
		ni, nj := i+dx[k], j+dy[k]
		if ni > -1 && ni < n && nj > -1 && nj < m {
			tot = Max(tot, dfs(grid, ni, nj))
		}
	}
	grid[i][j] = gold
	return tot + gold
}
