package main

/*
https://leetcode.cn/problems/flip-equivalent-binary-trees/
951. 翻转等价二叉树
*/
func flipEquiv(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && ((flipEquiv(a.Left, b.Left) && flipEquiv(a.Right, b.Right)) ||
		(flipEquiv(a.Left, b.Right) && flipEquiv(a.Right, b.Left)))
}
