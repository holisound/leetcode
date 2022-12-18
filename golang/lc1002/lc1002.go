package main

import (
	"math"
	"strings"

	. "github.com/holisound/leetcode/utils"
)

func commonChars(A []string) []string {
	/*
		这题我学到了:
		合理的使用packages, 这里使用strings.Count
	*/
	res := []string{}
	seen := map[rune]bool{}
	for _, k := range A[0] {
		if seen[k] {
			continue
		}
		seen[k] = true
		n := math.MaxInt32
		for _, w := range A {
			n = Min(n, strings.Count(w, string(k)))
		}
		for i := 0; i < n; i++ {
			res = append(res, string(k))
		}
	}
	return res
}
