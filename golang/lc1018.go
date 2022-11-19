func prefixesDivBy5(A []int) []bool {
	/*
		这题我学到了:
		取模,防止整型溢出
	*/
	cur := 0
	res := make([]bool, len(A))
	for i := range A {
		cur += A[i]
		res[i] = cur%5 == 0
		cur %= 5
		cur <<= 1
	}
	return res
}