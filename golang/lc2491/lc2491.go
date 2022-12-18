package main

/*
https://leetcode.cn/problems/divide-players-into-teams-of-equal-skill/
2491. 划分技能点相等的团队
*/
func dividePlayers(skill []int) int64 {
	tot := 0
	cnt := map[int]int{}
	for _, s := range skill {
		tot += s
		cnt[s]++
	}
	n := len(skill)
	each := tot / (n / 2)
	res := 0
	for _, s := range skill {
		if cnt[s] == 0 {
			continue
		}
		cnt[s]--
		if cnt[each-s] > 0 {
			cnt[each-s]--
			res += (each - s) * s
		} else {
			return -1
		}
	}
	return int64(res)
}
