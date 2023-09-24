package main

import "strings"

/*
https://leetcode.cn/problems/merge-strings-alternately/?envType=study-plan-v2&envId=leetcode-75
*/
func mergeAlternately(word1 string, word2 string) string {

	strb := new(strings.Builder)
	for i := 0; i < len(word1) && i < len(word2); i++ {
		strb.WriteByte(word1[i])
		strb.WriteByte(word2[i])
	}
	if len(word1) > len(word2) {
		word1, word2 = word2, word1
	}
	strb.WriteString(word2[len(word1):])
	return strb.String()
}
