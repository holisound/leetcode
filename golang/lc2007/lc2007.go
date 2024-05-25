package main

import "sort"

func findOriginalArray(changed []int) []int {
	sort.Ints(changed)
	ans := []int{}
	prefix := map[int]int{}
	for _, x := range changed {
		if x&1 == 0 && prefix[x/2] > 0 {
			prefix[x/2]--
			ans = append(ans, x/2)
			continue
		}
		prefix[x]++
	}
	if len(ans)*2 != len(changed) {
		return []int{}
	}
	return ans
}
