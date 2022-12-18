package main

var strs []string

func numDecodings(s string) int {
	/*
		这题我学到了：
		从递归到动态规划的算法演变过程
	*/
	// strs = []string{}
	// DFS TLE
	// dfs(s, 0)
	// return len(strs)
	return dpSolution(s)
}

func dfs(s string, start int) {
	if start == len(s) {
		strs = append(strs, s)
		return
	}
	if s[start] != '0' {
		dfs(s, start+1)
	}
	if start < len(s)-1 {
		if s[start] == '1' || s[start] == '2' && s[start+1] <= '6' {
			dfs(s, start+2)
		}
	}
}

func dpSolution(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 0; i < len(s); i++ {
		if s[i] != '0' {
			dp[i+1] += dp[i]
		}
		if i < len(s)-1 {
			if s[i] == '1' || s[i] == '2' && s[i+1] <= '6' {
				dp[i+2] += dp[i]
			}
		}
	}
	return dp[len(dp)-1]
}
