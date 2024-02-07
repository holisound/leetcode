package lc0221

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	area := 0
	cur, pre := make([]int, n), make([]int, n)
	for i := 0; i < m; i++ {
		cur, pre = pre, cur
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				cur[j] = 0
				continue
			}

			if i == 0 || j == 0 {
				cur[j] = 1
			} else {
				cur[j] = min(cur[j-1], pre[j-1], pre[j]) + 1
			}
			area = max(area, cur[j]*cur[j])
		}
	}

	return area
}
