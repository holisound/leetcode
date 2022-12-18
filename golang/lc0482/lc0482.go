package main

import (
	"strings"

	. "github.com/holisound/leetcode/utils"
)

func licenseKeyFormatting(S string, K int) string {
	/*
		这题我学到了；
		字符串取模拼接的一些技巧
	*/
	chars := []byte{}
	S = strings.ToUpper(S)
	for j := len(S) - 1; j > -1; j-- {
		if S[j] != '-' {
			if len(chars)%(K+1) == K { //
				chars = append(chars, '-')
			}
			chars = append(chars, S[j])
		}
	}
	Reverse(chars)
	return string(chars)

}
