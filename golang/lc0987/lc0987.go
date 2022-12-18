package main

import (
	"sort"

	. "github.com/holisound/leetcode/utils"
)

type Node struct {
	Val, X, Y int
}

func verticalTraversal(root *TreeNode) [][]int {
	/*
	   这题我学到了：
	   合理的使用结构体
	*/
	nodelist := []*Node{}
	dfs(root, 0, 0, &nodelist)

	sort.Slice(nodelist, func(i, j int) bool {
		n1, n2 := nodelist[i], nodelist[j]
		if n1.X == n2.X {
			if n1.Y == n2.Y {
				return n1.Val < n2.Val
			}
			return n1.Y < n2.Y
		}
		return n1.X < n2.X
	})
	res := [][]int{}
	if len(nodelist) < 1 {
		return res
	}
	res = append(res, []int{nodelist[0].Val})
	for i := 1; i < len(nodelist); i++ {
		pre, cur := nodelist[i-1], nodelist[i]
		if pre.X == cur.X {
			n := len(res)
			res[n-1] = append(res[n-1], cur.Val)
		} else {
			res = append(res, []int{cur.Val})
		}
	}
	return res
}

func dfs(root *TreeNode, x, y int, res *[]*Node) {
	if root == nil {
		return
	}
	*res = append(*res, &Node{root.Val, x, y})

	dfs(root.Left, x-1, y+1, res)
	dfs(root.Right, x+1, y+1, res)

}
