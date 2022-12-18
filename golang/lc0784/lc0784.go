package main

import "unicode"

func letterCasePermutation(S string) []string {
	/*
		这题我学到了：
		通过位运算来改变字母大小写
	*/
	res := []string{S}
	seen := map[string]bool{S: true}

	for i := range S {
		temp := []string{}
		for _, s := range res {
			chars := []rune(s)
			if unicode.IsLetter(chars[i]) {
				chars[i] ^= 32
			}
			t := string(chars)
			if !seen[t] {
				seen[t] = true
				temp = append(temp, t)
			}
		}
		res = append(res, temp...)
	}
	return res
}
