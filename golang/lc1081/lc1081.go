package main

import (
	"strings"

	. "github.com/holisound/leetcode/utils"
)

/*
https://leetcode.cn/problems/smallest-subsequence-of-distinct-characters/
1081. 不同字符的最小子序列
*/
func smallestSubsequence(s string) string {
	st := &Stack{}
	for i, c := range s {
		k := int(c)
		if st.Has(k) {
			continue
		}
		for st.Size() > 0 && st.Peek() > k &&
			strings.IndexRune(s[i+1:], rune(st.Peek())) != -1 {
			st.Pop()
		}

		st.Push(k)
	}
	sb := &strings.Builder{}
	for st.Size() > 0 {
		sb.WriteRune(rune(st.Pop()))
	}
	return reverseString(sb.String())
}

func reverseString(s string) string {
	chars := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
