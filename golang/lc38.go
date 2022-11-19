func countAndSay(n int) string {
    /*
       这题我学到了：
       byte跟int一样是可以自增的, 在计数并以字符串输出的时候非常方便高效
       无需使用strconv.Itoa
    */
    res := "1"
    for i := 1; i < n; i++ {
        var count byte = '1'
        temp := new(strings.Builder)
        for j := 1; j <= len(res); j++ {
            if j == len(res) || res[j-1] != res[j] {
                temp.WriteByte(count)
                temp.WriteByte(res[j-1])
                count = '1'
            } else {
                count++
            }
        }
        res = temp.String()
    }

    return res

}
