func orangesRotting(grid [][]int) int {
	que := list.New()
	cnt := 0
	ans := 0
	m, n := len(grid), len(grid[0])
	for i, row := range grid {
		for j, v := range row {
			if v == 2 {
				// 坏橘子\U0001f34a索引放队列里，用于扩散
				que.PushBack(&Pos{i, j, 0})
			} else if v == 1 {
				// 计数总共多少个好的橘子\U0001f34a
				cnt++
			}
		}
	}

	for que.Len() > 0 {
		u := que.Front()
		que.Remove(u)
		pos := u.Value.(*Pos)
		for _, neb := range getNeighbours(pos, m, n) {
			// 封装了getNeighbours，BFS相当常用
			if grid[neb.r][neb.c] == 1 {
				// 好橘子\U0001f34a被污染了
				grid[neb.r][neb.c] = 2
				cnt--
				que.PushBack(neb)
				ans = max(ans, neb.value)
			}
		}
	}
	if cnt != 0 {
		return -1
	}
	return ans
}
