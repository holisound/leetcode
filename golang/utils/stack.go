package utils

// 栈, 将int作为中间数据类型，在rune或byte之间转换

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
