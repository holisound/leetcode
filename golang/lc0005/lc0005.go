package main

/*
https://leetcode.cn/problems/longest-palindromic-substring/
5. 最长回文子串
*/
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	most, start, end := 0, 0, 0
	for r := 1; r < n; r++ {
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1] == 1) {
				dp[l][r] = 1
				if r-l+1 > most {
					most = r - l + 1
					start, end = l, r
				}
			}
		}
	}
	return s[start : end+1]
}
