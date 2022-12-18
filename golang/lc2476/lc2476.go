package main

import (
	"sort"

	. "github.com/holisound/leetcode/utils"
)

/*
https://leetcode.cn/problems/closest-nodes-queries-in-a-binary-search-tree/
2476. 二叉搜索树最近节点查询
*/

func closestNodes(root *TreeNode, queries []int) [][]int {
	bt := &BinaryTree{Root: root}
	nums := bt.InOrder()
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
