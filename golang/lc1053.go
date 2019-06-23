func prevPermOpt1(A []int) []int {
	index := -1
	for i := len(A) - 2; i > -1; i-- {
		if A[i] > A[i+1] {
			index = i
			break
		}
	}
	if index != -1 {
		moreBigger := A[index+1]
		swapIndex := index + 1
		for i := index + 2; i < len(A); i++ {
			if A[i] > moreBigger && A[i] < A[index] {
				moreBigger, swapIndex = A[i], i
			}
		}
		A[index], A[swapIndex] = A[swapIndex], A[index]

	}
	return A
}