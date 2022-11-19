func findNthDigit(n int) int {
	/*
		这题我学到了：
		整数位数计数的技巧
	*/
	base, d := 9, 1
	count := base * d
	num := 9
	for count < n {
		base *= 10
		d++
		if base*d+count > n {
			break
		}
		count += base * d
		num += base
	}
	rem := n - count    // 剩余总位数
	temp := num + rem/d // 数字取值
	if rem%d == 0 {     // 剩余0，取最后一位
		return temp % 10
	}
	temp++ // 剩余大于0, 取下一个数
	res := temp % 10
	for i := 1; i <= d-rem%d; i++ {
		temp /= 10
		res = temp % 10
	}
	return res
}