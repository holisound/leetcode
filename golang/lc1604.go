/*
https://leetcode.cn/problems/alert-using-same-key-card-three-or-more-times-in-a-one-hour-period/
1604. 警告一小时内使用相同员工卡大于等于三次的人
groupby key，暴力解法
*/
func alertNames(keyName []string, keyTime []string) []string {
    key2Times := map[string][]int{}
    for i, time := range keyTime {
        name := keyName[i]
        parts := strings.Split(time, ":")
        h, _ := strconv.Atoi(parts[0])
        m, _ := strconv.Atoi(parts[1])
        key2Times[name] = append(key2Times[name], h*60+m)
    }
    res := []string{}
    for name, times := range key2Times {
        for _, x := range times {
            cnt := 0
            for _, y := range times {
                if x > y || y > x+60 {
                    continue
                }
                cnt++
                if cnt == 3 {
                    goto flag
                }
            }
        }
        continue
        flag:res = append(res, name)
    }
    sort.Strings(res)
    return res
}
