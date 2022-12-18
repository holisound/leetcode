package utils

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 栈, 将int作为中间数据类型，在rune或byte之间转换
type Stack struct {
	data    []int
	inStack map[int]int
	sizeCur int
	sizeMax int
}

func (s *Stack) Push(x int) {
	s.sizeCur++
	if s.inStack == nil {
		s.inStack = make(map[int]int)
	}
	if s.sizeCur > s.sizeMax {
		s.data = append(s.data, x)
		s.sizeMax = s.sizeCur

	} else {
		s.data[s.sizeCur-1] = x
	}
	s.inStack[x]++
}

func (s *Stack) Pop() int {
	res := s.data[s.sizeCur-1]
	s.sizeCur--
	s.inStack[res]--
	return res
}

func (s *Stack) Has(x int) bool {
	return s.inStack[x] > 0
}

func (s *Stack) Size() int {
	return s.sizeCur
}

func (s *Stack) Peek() int {
	return s.data[s.sizeCur-1]
}

func Max(a ...int) int {
	x := a[0]
	for _, n := range a {
		if n > x {
			x = n
		}
	}
	return x
}

func Min(a ...int) int {
	x := a[0]
	for _, n := range a {
		if n < x {
			x = n
		}
	}
	return x
}

func Reverse(bits []byte) {
	for i, j := 0, len(bits)-1; i < j; i, j = i+1, j-1 {
		bits[i], bits[j] = bits[j], bits[i]
	}
}

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

type Graph struct {
}

func (g *Graph) Build(n int)
