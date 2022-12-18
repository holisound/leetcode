package main

func isToeplitzMatrix(matrix [][]int) bool {
	/*
		这题我学到了：
		矩阵比较运算的基本技巧
	*/
	for r := 0; r < len(matrix)-1; r++ {
		for c := 0; c < len(matrix[r])-1; c++ {
			if matrix[r][c] != matrix[r+1][c+1] {
				return false
			}
		}
	}
	return true
}
