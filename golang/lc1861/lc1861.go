package lc1861

func rotateTheBox(box [][]byte) [][]byte {
	m, n := len(box), len(box[0])
	ans := make([][]byte, n)
	for i := range ans {
		ans[i] = make([]byte, m)
		for j := range ans[i] {
			ans[i][j] = '.'
		}
	}
	for i := 0; i < m; i++ {
		stone := 0
		// 坐标转置 i, j -> j, m-1-i
		col := m - 1 - i
		for j := 0; j <= n; j++ {
			if j == n || box[i][j] == '*' {
				if j != n {
					ans[j][col] = box[i][j]
				}
				for k := j - 1; stone > 0; stone, k = stone-1, k-1 {
					ans[k][col] = '#'
				}
			} else if box[i][j] == '#' {
				stone++
			}
		}

	}
	return ans
}
