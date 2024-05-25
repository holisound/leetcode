package lc0498

func findDiagonalOrder(matrix [][]int) []int {
	dr, dc := -1, 1
	if len(matrix) == 0 {
		return []int{}
	}
	R, C := len(matrix), len(matrix[0])

	r, c := 0, 0

	res := make([]int, R*C)
	i := 0
	d := 1
	for i < R*C {
		if r >= 0 && r < R && c >= 0 && c < C {
			res[i] = matrix[r][c]
			i++
		}
		for r+dr >= 0 && r+dr < R && c+dc >= 0 && c+dc < C {
			r, c = r+dr, c+dc
			res[i] = matrix[r][c]
			i++
		}
		dr, dc = -dr, -dc
		if d == 1 {
			c++
		} else {
			r++
		}
		d = -d

	}
	return res
}
