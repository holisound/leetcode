// package lc1888

func minFlips(s string) int {
	n := len(s)
	cnt := 0 // 计数：与目标交替二进制字符串差异
	t := "10" // "01" or "10"
	for i := range s {
		if s[i] != t[i&1] {
			cnt++
		}
	}
	ans := min(cnt, n-cnt) // "01" or "10" 二元归一
	for i := range s {
		// 滑动窗口 + 索引处理
		if t[(i+n)&1] != s[i] {
			cnt++
		}
		if t[i&1] != s[i] {
			cnt--
		}
		ans = min(ans, cnt, n-cnt)
	}
	return ans
}
