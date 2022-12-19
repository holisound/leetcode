package main

import (
	. "github.com/holisound/leetcode/utils"
)

/*
https://leetcode.cn/problems/find-if-path-exists-in-graph/
1971. 寻找图中是否存在路径
*/
func validPath(n int, edges [][]int, source int, destination int) bool {
	uf := NewUnionFind(n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		uf.Merge(u, v)
	}
	return uf.Connect(source, destination)
}
