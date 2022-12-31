package lc2512

import (
	"sort"
	"strings"
)

/*
https://leetcode.cn/problems/reward-top-k-students/
2512. 奖励最顶尖的 K 名学生
*/
func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	id2points := map[int]int{}
	neg := map[string]bool{}
	for _, fb := range negative_feedback {
		neg[fb] = true
	}
	pos := map[string]bool{}
	for _, fb := range positive_feedback {
		pos[fb] = true
	}

	for i, r := range report {
		p := 0
		for _, w := range strings.Fields(r) {
			if neg[w] {
				p--
			} else if pos[w] {
				p += 3
			}
		}
		id2points[student_id[i]] = p
	}
	sort.Slice(student_id, func(i, j int) bool {
		if id2points[student_id[i]] == id2points[student_id[j]] {
			return student_id[i] < student_id[j]
		}
		return id2points[student_id[i]] > id2points[student_id[j]]
	})
	return student_id[:k]
}
