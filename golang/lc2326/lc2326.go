package lc2326

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(m int, n int, head *ListNode) [][]int {
	ans := make([][]int, m)
	for r := range ans {
		ans[r] = make([]int, n)
		for c := 0; c < n; c++ {
			ans[r][c] = -1
		}
	}
	r, c, dr, dc := 0, 0, 0, 1
	for head != nil {
		ans[r][c] = head.Val
		if nr, nc := r+dr, c+dc; nr <= -1 || nc <= -1 || nr >= m || nc >= n || ans[nr][nc] != -1 {
			dr, dc = dc, dr*-1
		}
		r, c = r+dr, c+dc
		head = head.Next
	}
	return ans
}
