func checkPossibility(nums []int) bool {
    /*
       这题我学到了：
       暴力解法：寻找第一次出现 nums[i-1] > nums[i] 的index (i-1, i)
       然后，有2种情况，使用nums[i] 改变 nums[i-1] 或 nums[i-1] 改变 nums[i]
       最后，两种情况分别遍历一次, 如果任意一种情况，不再出现nums[i-1] > nums[i]，
       则结果为true, 否则为false
    */
    if len(nums) <= 2 || sort.IntsAreSorted(nums) {
        return true
    }
    index := map[int]int{}
    for i := 1; i < len(nums); i++ {
        if nums[i-1] > nums[i] {
            index[i], index[i-1] = i-1, i
            break
        }
    }
    for k, v := range index {
        flag := false
        for i := 1; i < len(nums); i++ {
            i1, i2 := i-1, i
            if i1 == k {
                i1 = v
            } else if i2 == k {
                i2 = v
            }
            if nums[i1] > nums[i2] {
                flag = true
                break
            }

        }
        if !flag {
            return true
        }
    }
    return false

}
