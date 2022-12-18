package main

import (
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	digs, alph := []string{}, []string{}
	for _, log := range logs {
		j := strings.Index(log, " ")
		if unicode.IsLetter(rune(log[j+1])) {
			alph = append(alph, log)
		} else {
			digs = append(digs, log)
		}
	}

	sort.Slice(alph, func(i, j int) bool {
		// 在排序这里踩了点坑
		s1, s2 := alph[i], alph[j]
		p1, p2 := strings.Index(s1, " "), strings.Index(s2, " ")
		if s1[p1+1:] == s2[p2+1:] {
			return s1[:p1] < s2[:p2]
		}
		return s1[p1+1:] < s2[p2+1:]
	})

	return append(alph, digs...)
}
