package main

import (
	"unicode"

	"github.com/holisound/leetcode/utils"
)

func decodeString(s string) string {
	var alp []int
	stack := utils.CreateStack()
	for _, c := range s {
		if c == ']' {
			for stack.Size() != 0 && unicode.IsLetter(rune(stack.Peek())) {
				alp = append(alp, stack.Pop())
			}
			stack.Pop()
			num := 0
			base := 1
			for stack.Size() != 0 && unicode.IsDigit(rune(stack.Peek())) {
				d := int(rune(stack.Pop()) - '0')
				num += d * base
				base *= 10
			}
			for i := 0; i < num; i++ {
				for j := len(alp) - 1; j > -1; j-- {
					stack.Push(alp[j])
				}
			}
			alp = []int{}
			continue
		}
		stack.Push(int(c))
	}

	res := ""
	for stack.Size() != 0 {
		res = string(rune(stack.Pop())) + res
	}
	return res
}
