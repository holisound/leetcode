func longestArithSeqLength(A []int) int {
    /*
       这题我学到了<_>
    */
    dp := []map[int]int{}
    for i := 0; i < len(A); i++ {
        dp = append(dp, map[int]int{})
    }
    res := -1
    for i := 1; i < len(A); i++ {
        for j := 0; j < i; j++ {
            diff := A[i] - A[j]
            // j在前，i在后，判断在j位置是否出现过差值为diff的等差数列
            if num, ok := dp[j][diff]; ok {
                dp[i][diff] = num + 1
            } else {
                dp[i][diff] = 2
            }
            if dp[i][diff] > res {
                res = dp[i][diff]
            }
        }
    }
    return res
}
