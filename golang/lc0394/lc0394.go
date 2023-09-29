package main

import (
	"unicode"
)

func decodeString(s string) string {
	var stack, dig, alp []rune

	for _, c := range s {
		if c == ']' {
			for len(stack) != 0 && unicode.IsLetter(stack[len(stack)-1]) {
				alp = append(alp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			for len(stack) != 0 && unicode.IsDigit(stack[len(stack)-1]) {
				dig = append(dig, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			num := 0
			for i := len(dig) - 1; i > -1; i-- {
				num = num*10 + int(dig[i]-'0')
			}
			for i := 0; i < num; i++ {
				for j := len(alp) - 1; j > -1; j-- {
					stack = append(stack, alp[j])
				}
			}
			dig, alp = []rune{}, []rune{}
			continue
		}
		stack = append(stack, c)
	}
	return string(stack)
}
