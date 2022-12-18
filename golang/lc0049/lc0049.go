package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := map[string]int{}
	res := [][]string{}
	for _, str := range strs {
		chars := []rune(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		key := string(chars)
		if index, ok := m[key]; ok {
			res[index] = append(res[index], str)
		} else {
			m[key] = len(res)
			res = append(res, []string{str})
		}
	}
	return res
}
