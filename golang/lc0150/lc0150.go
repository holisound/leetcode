package main

import (
	"strconv"
	"strings"
)

func toInt(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func evalRPN(tokens []string) int {
	stack := make([]string, len(tokens))
	index := -1
	for _, t := range tokens {
		if strings.Index("+-*/", t) >= 0 {
			index--
			a, b := stack[index], stack[index+1]
			stack[index] = compute(toInt(a), toInt(b), t)
			continue
		}
		index++
		stack[index] = t
	}
	return toInt(stack[0])
}

var toStr = strconv.Itoa

func compute(x, y int, t string) string {
	if t == "+" {
		return toStr(x + y)
	} else if t == "-" {
		return toStr(x - y)
	} else if t == "*" {
		return toStr(x * y)
	}
	return toStr(x / y)
}
