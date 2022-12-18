package main

/*
https://leetcode.cn/problems/convert-to-base-2/
1017. 负二进制转换
*/
func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	table := []string{"0", "1"}
	res := ""
	for n != 0 {
		res = table[n&1] + res
		n >>= 1 // 当n < 0 不等价 n /= 2
		n *= -1
	}

	return res
}
