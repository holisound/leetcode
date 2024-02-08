package lc0115

func numDistinct(s string, t string) int {
	n := len(t)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := range s {
		for j := n - 1; j > -1; j-- {
			if s[i] == t[j] {
				dp[j+1] += dp[j]
			}
		}
	}
	return dp[n] % (1e9 + 7)
}
