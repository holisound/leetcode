package lc1926

import (
	"container/list"

	. "github.com/holisound/leetcode/utils"
)

type BFSUtilExt struct {
	BFSUtil
}

func (bfu *BFSUtilExt) Validate(pos *Pos) bool {
	return bfu.IsFirstVisited(pos) && (bfu.Matrix.([][]byte))[pos.R][pos.C] == '.'
}
func NewBFSUtil(matrix [][]byte) BFSUtilExt {
	nrow, ncol := len(matrix), len(matrix[0])
	visited := make([][]int, nrow)
	for i := range visited {
		visited[i] = make([]int, ncol)
	}
	return BFSUtilExt{BFSUtil{R: nrow, C: ncol, Visited: visited, Matrix: matrix}}

}
func nearestExit(maze [][]byte, entrance []int) int {
	// 1、确定BFS起点；如果起点为1个，可以扭转 起点 为 终点
	// 2、为了防止回到起点、终点不能和起点是同一个
	// 3、点的要素包括：index & value
	// 4、记录BFS迭代轮次：保存到Pos.value 或者 que.Len() 来 控制
	que := list.New()
	// que.PushBack(&Pos{entrance[0], entrance[1]})

	maze[entrance[0]][entrance[1]] = '+'
	m, n := len(maze), len(maze[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				if maze[i][j] == '.' {
					que.PushBack(&Pos{i, j, 0})
				}
			}
		}
	}

	bfu := NewBFSUtil(maze)

	for que.Len() != 0 {

		u := que.Front()
		que.Remove(u)
		pos := u.Value.(*Pos)

		bfu.SetValueByte(pos, '+')
		for _, neb := range bfu.GetNeighbours(pos) {
			if neb.R == entrance[0] && neb.C == entrance[1] {
				return neb.Value
			}
			if bfu.Validate(neb) {
				// bfu.SetValueByte(neb, '+') // 必须先标记 这些访问过的邻居节点，避免重复入队列
				que.PushBack(neb)
			}
		}

	}
	return -1
}
