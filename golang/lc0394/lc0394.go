package main

import (
	"unicode"

	lls "github.com/emirpasic/gods/stacks/linkedliststack"
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

func decodeStringV2(s string) string {
	stack := lls.New()
	var chars []rune
	for _, c := range s {
		if c == ']' {
			for val, ok := stack.Peek(); ok && unicode.IsLetter(val.(rune)); val, ok = stack.Peek() {
				chars = append(chars, val.(rune))
				stack.Pop()
			}
			stack.Pop()
			base, num := 1, 0
			for val, ok := stack.Peek(); ok && unicode.IsDigit(val.(rune)); val, ok = stack.Peek() {
				num, base = num+base*(int(val.(rune)-'0')), base*10
				stack.Pop()
			}
			for i := 0; i < num; i++ {
				for j := len(chars) - 1; j > -1; j-- {
					stack.Push(chars[j])
				}
			}
			chars = []rune{}
			continue
		}
		stack.Push(c)
	}
	res := ""
	for val, ok := stack.Peek(); ok; val, ok = stack.Peek() {
		res = string(val.(rune)) + res
		stack.Pop()
	}
	return res
}
