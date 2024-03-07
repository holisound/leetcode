package lc3034

func countMatchingSubarrays(nums []int, pattern []int) int {
	n, m := len(nums), len(pattern)
	p := make([]int, n-1)
	ans := 0
	for i := 1; i < n; i++ {
		if nums[i-1] < nums[i] {
			p[i-1] = 1
		} else if nums[i-1] > nums[i] {
			p[i-1] = -1
		}
		cnt := 0
		for j, k := i-m, 0; k < m && j > -1 && j+1 < n; k, j = k+1, j+1 {
			if pattern[k] == p[j] {
				cnt++
			} else {
				break
			}
		}
		if cnt == m {
			ans++
		}
	}
	return ans
}
