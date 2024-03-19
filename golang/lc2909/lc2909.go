package lc2909

func minimumSum(nums []int) int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		left[i] = math.MaxInt
		right[i] = math.MaxInt
	}
	left[0], right[n-1] = nums[0], nums[n-1]
	for i := 1; i < n; i++ {
		left[i] = min(nums[i], left[i-1])
		right[n-i-1] = min(nums[n-i-1], right[n-i])
	}

	ans := math.MaxInt
	for i := 1; i+1 < n; i++ {
		if left[i-1] < nums[i] && nums[i] > right[i+1] {
			ans = min(ans, left[i-1]+nums[i]+right[i+1])
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
