var ans int

func longestZigZag(root *TreeNode) int {
	ans = 0
	postOrder(root)
	return ans
}

func postOrder(root *TreeNode) []int {
	if root == nil {
		return []int{-1, -1}
	}
	l, r := postOrder(root.Left), postOrder(root.Right)
	ans = max(ans, l[1]+1, r[0]+1)
	return []int{
		l[1] + 1, r[0] + 1,
	}
}
