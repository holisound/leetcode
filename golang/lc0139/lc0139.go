package lc0139

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 0; i <= n; i++ {
		for _, w := range wordDict {
			if k := len(w); i-k >= 0 && s[i-k:i] == w && dp[i-k] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
