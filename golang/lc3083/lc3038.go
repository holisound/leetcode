package lc3083

func isSubstringPresent(s string) bool {
	// 检查相邻的2个元素
	// 形如 aa, aba, ab...ba 都是有效子串
	exists := map[string]bool{}
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] || exists[s[i-1:i+1]] {
			return true
		}
		exists[string([]byte{s[i], s[i-1]})] = true
	}
	return false
}
