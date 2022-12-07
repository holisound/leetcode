/*
https://leetcode.cn/problems/amount-of-time-for-binary-tree-to-be-infected/
2385. 感染二叉树需要的总时间
*/
var (
    res int
    graph map[int][]int
    vis map[int]bool
)
func buildGraph(root *TreeNode) {
    if root != nil {
        for _, node := range []*TreeNode{root.Left, root.Right} {
            if node != nil {
                graph[root.Val] = append(graph[root.Val], node.Val)
                graph[node.Val] = append(graph[node.Val], root.Val)
            }
        }
        buildGraph(root.Left)
        buildGraph(root.Right)
    }
}
func traverse(u, depth int) {
    if vis[u] {
        return
    }
    vis[u] = true
    if depth > res {
        res = depth
    }
    for _, v := range graph[u] {
        traverse(v, depth+1)
    }

}
func amountOfTime(root *TreeNode, start int) int {
    res = 0
    graph = map[int][]int{}
    vis = map[int]bool{}
    buildGraph(root)
    traverse(start,0)
    return res
}
