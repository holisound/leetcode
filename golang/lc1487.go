/*
https://leetcode.cn/problems/making-file-names-unique/
1487. 保证文件名唯一
*/ 
func getFolderNames(names []string) []string {
    res := []string{}
    cnt := map[string]int{}
    
    for _, name := range names {
        i,k:=cnt[name],name
        for ;cnt[k]>0;i,k=i+1,fmt.Sprintf("%s(%d)", name, i){} 
        cnt[name] = i
        cnt[k]=1
        res = append(res, k)
    }

    return res
}
