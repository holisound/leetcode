package lc3076

func shortestSubstrings(arr []string) []string {
	n := len(arr)
	ans := make([]string, n)
	m := map[string]int{}
	for _, s := range arr {
		// substr(s,i,j)
		seen := map[string]bool{} // 同一字符串的相同子串只计数一次
		for j := 0; j < len(s); j++ {
			for i := 0; i <= j; i++ {
				if p := s[i : j+1]; !seen[p] {
					seen[p] = true
					m[p]++
				}
			}
		}
	}
	for idx, s := range arr {
		cond := ""
		shortest := len(s) + 1
		for j := 0; j < len(s); j++ {
			for i := 0; i <= j; i++ {
				if m[s[i:j+1]] != 1 { // 不同字符串出现相同子串
					continue
				}
				if k := j - i + 1; k < shortest {
					shortest, cond = k, s[i:j+1]
				} else if k == shortest {
					cond = min(cond, s[i:j+1])
				}
			}
		}
		ans[idx] = cond
	}
	return ans
}
