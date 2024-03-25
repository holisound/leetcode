package lc2976

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	// 建图
	graph := [26][26]int{}
	for i := range cost {
		u, v, w := original[i]-'a', changed[i]-'a', cost[i]
		if graph[u][v] == 0 {
			graph[u][v] = w
		} else {
			graph[u][v] = min(graph[u][v], w)
		}
	}
	// 建图之后，如果graph[u][v]==0，则表示节点u,v之间不存在边，可能是通过其他中间节点连接的
	ans := int64(0)
	type Node struct {
		Id      int
		pathSum [26]int
		total   int
	}
	que := []*Node{}
	bestPath := [26][26]int{}
	for i := range source {
		if source[i] == target[i] {
			continue
		}
		u, v := int(source[i]-'a'), int(target[i]-'a')
		if bp := bestPath[u][v]; bp != 0 {
			ans += int64(bp)
			continue
		}
		node := &Node{u, [26]int{}, 0}
		que = append(que, node)
		best := math.MaxInt32
		for len(que) != 0 {
			node := que[0]
			u, pathSum := node.Id, node.pathSum[node.Id]
			que = que[1:]
			if u == v {
				best = min(best, pathSum)
				continue
			}
			for v, w := range graph[u] {
				if w != 0 && node.pathSum[v] == 0 && node.pathSum[u]+w < best {
					// 入队列前做好节点（种子）筛选
					pathSum := [26]int{}
					for i, x := range node.pathSum {
						pathSum[i] = x
					}
					pathSum[v] = pathSum[u] + w
					node := &Node{v, pathSum, pathSum[v]}
					que = append(que, node)
				}
			}
		}

		if best == math.MaxInt32 {
			return -1
		}
		bestPath[u][v] = best
		ans += int64(best)
	}

	return ans
}
