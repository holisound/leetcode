package main

import "sort"

/*
https://leetcode.cn/problems/closest-nodes-queries-in-a-binary-search-tree/
2476. 二叉搜索树最近节点查询
*/

var (
	nums []int
)

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		nums = append(nums, root.Val)
		dfs(root.Right)
	}
}
func closestNodes(root *TreeNode, queries []int) [][]int {
	nums = []int{}
	dfs(root)
	res := [][]int{}
	for _, q := range queries {
		pos := sort.SearchInts(nums, q)
		one := []int{-1, -1}
		if pos < len(nums) {
			one[1] = nums[pos]
			if nums[pos] > q && pos > 0 {
				one[0] = nums[pos-1]
			} else if nums[pos] == q {
				one[0] = nums[pos]
			}
		} else {
			one[0] = nums[len(nums)-1]
		}
		res = append(res, one)
	}
	return res
}
