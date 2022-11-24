/*
https://leetcode.cn/problems/expressive-words/
809. 情感丰富的文字
*/

func expressiveWords(s string, words []string) int {
    res :=0
    for _, w := range words {
        c1, c2 := count(s), count(w)
        if len(c1) != len(c2) {
            continue
        }
        flag := false
        for i := 0;i < len(c1);i++ {
            a, b := c1[i], c2[i]
            if a[0] != b[0] || a[1]<3&&b[1]<a[1] || a[1]<b[1]{
                flag = true
                break
            } 
        }
        if flag {
            continue
        }
        res++
    }

    return res
}

func count(s string) [][]int {
    dp := make([]int, len(s))
    for i:=range dp{
        dp[i]=1
    }
    cnt := [][]int{}
    for i:=1;i<=len(s);i++{
        if i < len(s) && s[i-1]==s[i]{
            dp[i]=dp[i-1]+1
        } else {
            cnt = append(cnt, []int{int(s[i-1]), dp[i-1]})
        }
    }
    return cnt
}
