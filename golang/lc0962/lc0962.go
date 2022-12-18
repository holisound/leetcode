package main

import (
	. "github.com/holisound/leetcode/utils"
)

/*
https://leetcode.cn/problems/maximum-width-ramp/
962. 最大宽度坡
关键字:单调栈
*/
func maxWidthRamp(nums []int) int {
	st := &Stack{}
	for i := range nums {
		if st.Size() == 0 || nums[st.Peek()] > nums[i] {
			st.Push(i)
		}
	}
	res := 0
	for i := len(nums) - 1; i > -1; i-- {
		for st.Size() > 0 && nums[st.Peek()] <= nums[i] {
			j := st.Pop()
			if i-j > res {
				res = i - j
			}
		}
	}
	return res
}
