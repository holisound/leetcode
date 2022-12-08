/*
https://leetcode.cn/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/
2316. 统计无向图中无法互相到达点对数
*/
var (
    graph map[int][]int
    vis map[int]bool
)
func dfs(u int) int {
    if vis[u] {
        return 0
    }
    vis[u]=true
    count := 1
    for _, v:=range graph[u] {
        count += dfs(v)
    }
    return count
}
func countPairs(n int, edges [][]int) int64 {
    graph = map[int][]int{}
    vis = map[int]bool{}
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        graph[u] = append(graph[u], v)
        graph[v] = append(graph[v], u)
    }
    var tot, ans int
    for u:=0;u<n;u++{
        num:=dfs(u)
        ans += num*tot
        tot+=num
    }
    return int64(ans)
}
