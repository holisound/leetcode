/*
https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree/
863. 二叉树中所有距离为 K 的结点
*/
var (
    res []int
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
func traverse(u, k int) {
    if vis[u] {
        return
    }
    vis[u] = true
    if k == 0 {
        res = append(res, u)
        return
    }
    for _, v := range graph[u] {
        traverse(v, k-1)
    }

}
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    res = []int{}
    vis = map[int]bool{}
    graph = map[int][]int{}
    buildGraph(root)
    traverse(target.Val, k)
    return res
}
