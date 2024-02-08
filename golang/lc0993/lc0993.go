package lc0993

func isCousins(root *TreeNode, x int, y int) bool {
	store := map[int]struct {
		depth, parent int
	}{}

	var dfs func(root *TreeNode, p, d int)
	dfs = func(root *TreeNode, p, d int) {
		if root != nil {
			store[root.Val] = struct {
				depth, parent int
			}{d, p}
			dfs(root.Left, root.Val, d+1)
			dfs(root.Right, root.Val, d+1)
		}
	}
	dfs(root, -1, 0)
	return store[x].parent != store[y].parent && store[x].depth == store[y].depth
}
