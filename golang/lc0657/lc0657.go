package main

func judgeCircle(moves string) bool {
	/*
		这题我学到了：
		初步形成解决坐标位移类问题的算法框架
	*/
	index := map[rune]int{
		'U': 0, 'D': 1, 'L': 2, 'R': 3,
	}
	dx := []int{0, 0, -1, 1}
	dy := []int{1, -1, 0, 0}
	x, y := 0, 0
	for _, m := range moves {
		x, y = x+dx[index[m]], y+dy[index[m]]
	}
	return x == 0 && y == 0
}
