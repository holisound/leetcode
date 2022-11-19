func checkPossibility(nums []int) bool {
    index := -1
    n := len(nums)
    for i := 1; i < n; i++ {
        if nums[i-1] > nums[i]{
            index = i-1
            break
        }
    }
    if index == -1 {
        return true
    }
    temp := nums[index]
    nums[index] = nums[index+1]
    if isAsc(nums){
        return true
    }
    nums[index], nums[index+1] = temp, temp
    if isAsc(nums){
        return true
    }
    return false
    
}

func isAsc(a []int)bool{
    for i := 1; i < len(a);i++ {
        if a[i-1] > a[i]{
            return false
        }
    }
    return true
}
