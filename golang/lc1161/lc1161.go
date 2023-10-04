package lc1161

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type TreeNodeExt struct {
	*TreeNode
	Index int
}

func maxLevelSum(root *TreeNode) int {
	stack := []*TreeNodeExt{{root, 0}}
	sum := []int{}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			stack = append(stack, &TreeNodeExt{node.Left, node.Index + 1})
		}
		if node.Right != nil {
			stack = append(stack, &TreeNodeExt{node.Right, node.Index + 1})
		}
		if len(sum) == node.Index {
			sum = append(sum, node.Val)
		} else {
			sum[node.Index] += node.Val
		}
	}
	// 因为节点TreeNode.Val存在负数，所以只能在遍历结束后对sumMax取值进行判断
	sumMax := max(sum...)
	for i := range sum {
		if sum[i] == sumMax {
			return i + 1
		}
	}
	return 0
}
func max(a ...int) int {
	x := a[0]
	for _, n := range a {
		if n > x {
			x = n
		}
	}
	return x
}
