package main

import (
	. "github.com/holisound/leetcode/utils"
)

func countSquares(matrix [][]int) int {
	res := 0
	for i, row := range matrix {
		for j, x := range row {
			if i > 0 && j > 0 && x == 1 {
				matrix[i][j] += Min(matrix[i-1][j-1], matrix[i][j-1], matrix[i-1][j])
			}
			res += matrix[i][j]
		}
	}
	return res
}
