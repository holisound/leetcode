package main

import (
	. "github.com/holisound/leetcode/utils"
)

/*
https://leetcode.cn/problems/minimum-number-of-swaps-to-make-the-binary-string-alternating/
1864. 构成交替字符串需要的最小交换次数
*/
func minSwaps(s string) int {
	/*
	   010101...
	   101010...
	*/
	var a, b, m, n int
	for _, c := range s {
		i := int(c - '0')
		if i == 0 {
			b++
		} else {
			b--
		}
		if i == a {
			m++
		} else if i == a^1 {
			n++
		}
		a ^= 1
	}
	if b < -1 || b > 1 {
		return -1
	}
	if b == -1 {
		// case#118 count({'0': 122, '1': 123})
		return (m + 1) / 2
	} else if b == 1 {
		return (n + 1) / 2
	}
	return Min(m, n) / 2
}
