package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
https://leetcode.cn/problems/minimize-result-by-adding-parentheses-to-expression/
2232. 向表达式添加括号后的最小结果
*/

func minimizeResult(expression string) string {
	n := len(expression)
	mid := strings.Index(expression, "+")
	least := math.MaxInt32
	var ans string
	for i := 0; i < mid; i++ {
		for j := mid + 1; j < n; j++ {
			a := 1
			if len(expression[:i]) != 0 {
				a, _ = strconv.Atoi(expression[:i])
			}
			b, _ := strconv.Atoi(expression[i:mid])
			c, _ := strconv.Atoi(expression[mid+1 : j+1])
			d := 1
			if len(expression[j+1:]) != 0 {
				d, _ = strconv.Atoi(expression[j+1:])
			}
			if a*(b+c)*d < least {
				least = a * (b + c) * d
				ans = fmt.Sprintf("%s(%s+%s)%s",
					expression[:i], expression[i:mid],
					expression[mid+1:j+1], expression[j+1:])
			}
		}
	}
	return ans
}
