package lc0199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type TreeNodeExt struct {
	*TreeNode
	Index int
}

func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	stack := []*TreeNodeExt{{root, 0}}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 正常来说是node.Right先入栈，这里取的是右视图，所以node.Left可以先入栈，简化代码
		if node.Left != nil {
			stack = append(stack, &TreeNodeExt{node.Left, node.Index + 1})
		}
		if node.Right != nil {
			stack = append(stack, &TreeNodeExt{node.Right, node.Index + 1})
		}
		if node.Index == len(ans) {
			ans = append(ans, node.Val)
		}
		// node.Right 先入栈的话，需要更新ans[i]
		// else {
		// 	ans[node.Index] = node.Val
		// }
	}

	return ans
}
