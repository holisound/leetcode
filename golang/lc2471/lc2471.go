package lc2471

func minimumOperations(root *TreeNode) int {

	levels := [][]int{}
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root != nil {
			if depth == len(levels) {
				levels = append(levels, []int{})
			}
			levels[depth] = append(levels[depth], root.Val)

			dfs(root.Left, depth+1)
			dfs(root.Right, depth+1)
		}
	}
	dfs(root, 0)
	ans := 0

	for _, level := range levels {

		n := len(level)
		cnt := 0
		for start := 0; start < n; start++ {
			minIndex := start
			for i := minIndex; i < n; i++ {
				if level[i] < level[minIndex] {
					minIndex = i
				}
			}
			if minIndex == start {
				continue
			}
			level[start], level[minIndex] = level[minIndex], level[start]
			cnt++
		}
		ans += cnt

	}
	// fmt.Println(levels)
	return ans
}
