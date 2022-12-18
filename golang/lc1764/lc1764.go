package main

/*
https://leetcode.cn/problems/form-array-by-concatenating-subarrays-of-another-array/
1764. 通过连接另一个数组的子数组得到一个数组
*/
func canChoose(groups [][]int, nums []int) bool {
	var start int
	for _, grp := range groups {
		var flag bool
		for i := start; i < len(nums)-len(grp)+1; i++ {
			flag = true
			for j, x := range grp {
				if x != nums[i+j] {
					flag = false
					break
				}
			}
			if flag {
				start = i + len(grp)
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}
