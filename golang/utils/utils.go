package utils

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
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

// 曼哈顿距离
func ManhattanDist(x1, y1, x2, y2 int) int {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	return x2 - x1 + y2 - y1
}
