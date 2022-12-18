package main

/*
https://leetcode.cn/problems/reachable-nodes-with-restrictions/
2368. 受限条件下可到达节点的数目
*/
var (
	graph map[int][]int
	vis   map[int]bool
	res   int
)

func dfs(u int) {
	if vis[u] {
		return
	}
	vis[u] = true
	res++
	for _, v := range graph[u] {
		dfs(v)
	}
}
func reachableNodes(n int, edges [][]int, restricted []int) int {
	graph = map[int][]int{}
	vis = map[int]bool{}
	for _, n := range restricted {
		vis[n] = true
	}
	res = 0
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	dfs(0)
	return res
}
