package utils

type BinaryTree struct {
	Root *TreeNode
}

func (bt *BinaryTree) InOrder() []int {
	values := []int{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			values = append(values, root.Val)
			dfs(root.Right)
		}
	}
	dfs(bt.Root)
	return values
}

func (bt *BinaryTree) BuildGraph() map[int][]int {
	graph := make(map[int][]int)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			for _, node := range []*TreeNode{root.Left, root.Right} {
				if node != nil {
					graph[root.Val] = append(graph[root.Val], node.Val)
					graph[node.Val] = append(graph[node.Val], root.Val)
				}
			}
			dfs(root.Left)
			dfs(root.Right)
		}
	}
	dfs(bt.Root)
	return graph
}

func (bt *BinaryTree) Traverse() {

}
