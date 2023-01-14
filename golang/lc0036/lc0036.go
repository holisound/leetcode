package lc0036

/*
https://leetcode.cn/problems/valid-sudoku/
36. Valid Sudoku
*/
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		row, col := map[byte]int{}, map[byte]int{}
		flag := true
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				row[board[i][j]]++
			}
			if board[j][i] != '.' {
				col[board[j][i]]++
			}
			if col[board[j][i]] > 1 || row[board[i][j]] > 1 {
				flag = false
				break
			}
		}
		if !flag {
			return false
		}
	}
	return search(board)
}

var dirs [][]int = [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}, {2, 1}, {2, 0}, {1, 0}, {1, 1}}

func search(board [][]byte) bool {
	for _, i := range []int{0, 3, 6} {
		for _, j := range []int{0, 3, 6} {
			box := map[byte]int{}
			for _, d := range dirs {
				ni, nj := i+d[0], j+d[1]
				if board[ni][nj] != '.' {
					box[board[ni][nj]]++
					if box[board[ni][nj]] > 1 {
						return false
					}
				}
			}
		}
	}
	return true
}
