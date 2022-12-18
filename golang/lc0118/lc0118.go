package main

func generate(numRows int) [][]int {
	/*
	   这题我学到了：
	   使用2种方法求解
	   1、递推
	   2、组合数
	*/
	res := make([][]int, numRows)
	for r := 0; r < numRows; r++ {
		res[r] = make([]int, r+1)
		// i, j := 0, len(res[r])-1
		// res[r][i] = 1
		// res[r][j] = 1
		// if r > 0 {
		//     for c := i+1; c < j; c++ {
		//         res[r][c] = res[r-1][c] + res[r-1][c-1]
		//     }
		// }
		for c := 0; c <= r; c++ {
			res[r][c] = Comb(r, c)
		}
	}

	return res
}

func Comb(n, k int) int {
	if k == 0 {
		return 1
	}
	if n < k {
		return 0
	}
	if n-k < k {
		k = n - k
	}
	res := 1
	for i := 0; i < k; i++ {
		res = res * (n - i) / (i + 1)
	}

	return res
}
