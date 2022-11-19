func smallestSubsequence(text string) string {
    /*
       这题我学到了：
        第一次遍历，先对字符计数
        第二次遍历，按顺序依次减少剩余字符的数量，同时利用Stack弹出非最优解
    */
    stack := make([]rune, len(text))
    endIndex := -1
    seen := map[rune]bool{}
    count := map[rune]int{}
    for _, c := range text {
        count[c]++
    }
    for _, c := range text {
        count[c]--
        if seen[c] {
            continue
        }
        for endIndex > -1 {
            v := stack[endIndex]
            if v > c && count[v] > 0 {
                endIndex--
                delete(seen, v)
            } else {
                break
            }
        }
        endIndex++
        seen[c] = true
        stack[endIndex] = c
    }
    return string(stack[:endIndex+1])

}
