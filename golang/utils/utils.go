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
