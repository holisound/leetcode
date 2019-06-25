func addNegabinary(arr1 []int, arr2 []int) []int {
    /*
       这题我学到了：
       1.在golang中，“%”运算的如果是负数，返回值也是负数，所以，这里应该采用“&”运算符
       2.从末尾开始, 当index为偶数，值是+正数，奇数，则为-负数
       arr1 = [1,1,1,1,1]
           arr2 = [1,0,1]
                       0, +2
                     0, -2+2 = 0
                   0, 4+4 = 8
                 0, -8 + 8
               1
    */
    l1, l2 := len(arr1), len(arr2)
    i, j, k, carry := l1-1, l2-1, l1+l2, 0
    res := make([]int, l1+l2+1)
    for i > -1 || j > -1 || carry != 0 {
        s := -carry //如果有上一次迭代有进位，应该抵消
        if i > -1 {
            s += arr1[i]
        }
        if j > -1 {
            s += arr2[j]
        }
        carry = s >> 1
        if s == -1 {
            // 进位未抵消
            carry = -1
        }
        res[k] = s & 1
        k--
        i, j = i-1, j-1
    }
    index := -1
    for i := 0; i < len(res); i++ {
        if res[i] == 1 {
            index = i
            break
        }
    }
    if index == -1 {
        return []int{0}
    }
    return res[index:]
}
