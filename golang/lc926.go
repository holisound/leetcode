/*
https://leetcode.cn/problems/flip-string-to-monotone-increasing/submissions/
926. 将字符串翻转到单调递增
思路：
1、将01看成2个部分
2、需要翻转次数：当前元素之前出现的”1“和当前元素之后出现的“0“数量之和最小值
*/ 
func minFlipsMonoIncr(s string) int {
    zeroTotal := strings.Count(s, "0")
    res := zeroTotal
    one, zero := 0, 0
    for _, c := range s {
        if c == '1' {
            one++
        } else {
            zero++
        }
        res = min(res, one+zeroTotal-zero)
    }
    return res
}

func min(a, b int) int {
    if a < b{
        return a
    }
    return b
}
