package main

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
