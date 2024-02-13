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
	var dfs func(root *TreeNode, x, y int)
	dfs = func(root *TreeNode, x, y int) {
		if root == nil {
			return
		}
		nodelist = append(nodelist, &Node{root.Val, x, y})
		dfs(root.Left, x-1, y+1)
		dfs(root.Right, x+1, y+1)
	}
	dfs(root, 0, 0)
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
	pre := nodelist[0]
	res = append(res, []int{})
	for _, cur := range nodelist {
		if pre.X == cur.X {
			n := len(res)
			res[n-1] = append(res[n-1], cur.Val)
		} else {
			res = append(res, []int{cur.Val})
		}
		pre = cur
	}
	return res
}
