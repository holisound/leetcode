func lenLongestFibSubseq(A []int) int {
	/*
		这题我学到了：
		图的一些基本概念：结点，路径
	*/
	res := 0
	index := map[int]int{}
	for i, a := range A {
		index[a] = i
	}
	longest := map[int]int{}
	N := len(A)
	for k, a := range A {
		for j := 0; j < k; j++ {
			if i, ok := index[a-A[j]]; ok && i < j {
				if n, ok := longest[N*i+j]; ok {
					longest[N*j+k] = n + 1
				} else {
					longest[N*j+k] = 3
				}
				if longest[N*j+k] > res {
					res = longest[N*j+k]
				}
			}
		}
	}

	return res
}