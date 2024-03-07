package lc3071

func minimumOperationsToWriteY(grid [][]int) int {
	n := len(grid)
	cntY := make([]int, 3)
	cntNY := make([]int, 3)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := grid[i][j]
			if i <= n/2 && (i == j || i == n-j-1) {
				cntY[x]++
			} else if i > n/2 && j == n/2 {
				cntY[x]++
			} else {
				cntNY[x]++
			}
		}
	}
	// 3->4, 5->7, 7->10
	// 字母Y的点总数：totY := 3*n+1
	ans := n * n
	for i, y := range cntY {
		for j, ny := range cntNY {
			if i != j {
				ans = min(ans, n*n-y-ny)
			}
		}
	}
	return ans
}
