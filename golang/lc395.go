func longestSubstring(s string, k int) int {
    /*
       这题我学到了：
       计数(countk), 迭代时统计符合数量(>=k)的字符数量
       和递归的技巧
    */
    count := map[rune]int{}
    countk := 0
    for _, c := range s {
        count[c]++
        if count[c] == k {
            countk++
        }
    }
    if countk == len(count) {
        return len(s)
    }
    j, res, n := 0, 0, len(s)
    for i := 0; i <= n; i++ {
        if i == n || count[rune(s[i])] < k {
            l := longestSubstring(s[j:i], k)
            if l > res {
                res = l
            }
            j = i + 1
        }
    }
    return res
}
