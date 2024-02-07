package main

/*
https://leetcode.cn/problems/longest-palindromic-substring/
5. 最长回文子串
*/
func longestPalindrome(s string) string {
	res := s[:1]
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if s[i] == s[j] && (i-j <= 2 || dp[j+1][i-1] == 1) {
				dp[j][i] = 1
				if i-j+1 > len(res) {
					res = s[j : i+1]
				}
			}
		}
	}

	return res
}
