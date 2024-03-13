package lc3067

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1  // 提示中有说明
	count := make([]int, n)
	// 建立节点与边长映射,用哈希表（不用矩阵会超内存）
	graph := make([]map[int]int, n)
	for i := range graph {
		graph[i] = make(map[int]int)
	}
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		graph[u][v] = w
		graph[v][u] = w
	}
	var dfs func(u, parent, pathSum int) int
	dfs = func(u, parent, pathSum int) int {
		cnt := 0
		if pathSum%signalSpeed == 0 {
			cnt = 1
		}
		for v, w := range graph[u] {
			if v != parent {
				cnt += dfs(v, u, pathSum+w)
			}
		}
		return cnt
	}
	for u := range count {
		s := 0
		for v, w := range graph[u] {
			cnt := dfs(v, u, w)
			count[u] += s * cnt
			s += cnt
		}
	}

	return count
}
