package main

import (
	"container/list"

	. "github.com/holisound/leetcode/utils"
)

type BFSUtilExt struct {
	BFSUtil
}

func (bfu *BFSUtilExt) Validate(pos *Pos) bool {
	return bfu.IsFirstVisited(pos) && (bfu.Matrix.([][]int))[pos.R][pos.C] == 1
}
func NewBFSUtil(matrix [][]int) BFSUtilExt {
	nrow, ncol := len(matrix), len(matrix[0])
	visited := make([][]int, nrow)
	for i := range visited {
		visited[i] = make([]int, ncol)
	}
	return BFSUtilExt{BFSUtil{R: nrow, C: ncol, Visited: visited, Matrix: matrix}}

}

func orangesRotting(grid [][]int) int {
	que := list.New()
	cnt := 0
	ans := 0

	for i, row := range grid {
		for j, v := range row {
			if v == 2 {
				// 坏橘子\U0001f34a索引放队列里，用于扩散
				que.PushBack(&Pos{i, j, 0})
			} else if v == 1 {
				// 计数总共多少个好的橘子\U0001f34a
				cnt++
			}
		}
	}
	bfu := NewBFSUtil(grid)
	for que.Len() > 0 {
		u := que.Front()
		que.Remove(u)
		pos := u.Value.(*Pos)
		for _, neb := range bfu.GetNeighbours(pos) {
			// 封装了getNeighbours，BFS相当常用
			if bfu.Validate(neb) {
				// 好橘子\U0001f34a被污染了
				cnt--
				que.PushBack(neb)
				if neb.Value > ans {
					ans = neb.Value
				}
			}
		}
	}
	if cnt != 0 {
		return -1
	}
	return ans
}
