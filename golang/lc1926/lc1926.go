func nearestExit(maze [][]byte, entrance []int) int {
	// 1、确定BFS起点；如果起点为1个，可以扭转 起点 为 终点
	// 2、为了防止回到起点、终点不能和起点是同一个
	// 3、点的要素包括：index & value
	// 4、记录BFS迭代轮次：保存到Pos.value 或者 que.Len() 来 控制
	que := list.New()
	maze[entrance[0]][entrance[1]] = '+'
	m, n := len(maze), len(maze[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				if maze[i][j] == '.' {
					que.PushBack(&Pos{i, j, 0})
				}
			}
		}
	}

	for que.Len() != 0 {
		u := que.Front()
		que.Remove(u)
		pos := u.Value.(*Pos)
		maze[pos.r][pos.c] = '+'
		for _, neb := range getNeighbours(pos, m, n) {
			if neb.r == entrance[0] && neb.c == entrance[1] {
				return neb.value
			}
			if maze[neb.r][neb.c] == '.' {
				maze[neb.r][neb.c] = '+' // 必须先标记 这些访问过的邻居节点，避免重复入队列
				que.PushBack(neb)
			}
		}
	}
	return -1
}
