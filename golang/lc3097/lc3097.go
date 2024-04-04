package lc3097

func minimumSubarrayLength(nums []int, k int) int {
	n := len(nums)
	ans := n + 1
	bitCnt := make([]int, 32)
	mask := 0
	j := 0
	for i, x := range nums {
		for t := 0; t < 32; t++ {
			bitCnt[t] += (x >> t) & 1
		}
		mask |= x
		for mask >= k && j <= i {
			ans = min(ans, i-j+1)
			for t := 0; t < 32; t++ {
				y := (nums[j] >> t) & 1
				bitCnt[t] -= y
				if y != 0 && bitCnt[t] == 0 {
					mask ^= 1 << t
				}
			}
			j++
		}
	}
	if ans == n+1 {
		return -1
	}
	return ans
}
