package lc1685

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	tot := 0
	for _, x := range nums {
		tot += x
	}

	ans := make([]int, n)
	leftSum := 0
	for i, x := range nums {
		rightSum := tot - x - leftSum
		// x=nums[i]有n-1个，因为|nums[i]-nums[j]|只有n-1对
        // 利用【单调非递减】特性，求解绝对值之和
        // nums[i]左侧k个nums[j]之和肯定小于等于k*nums[i]
        // nums[i]右侧k个nums[j]之和肯定大于等于k*nums[i]
		ans[i] = rightSum - x*(n-i-1) + x*i - leftSum
		leftSum += x
	}

	return ans
}
