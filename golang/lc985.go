func sumEvenAfterQueries(A []int, queries [][]int) []int {
	total := 0
	for _, a := range A {
		if a%2 == 0 {
			total += a
		}
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		val, index := q[0], q[1]
		if A[index]%2 == 0 {
			total -= A[index]
		}
		A[index] += val
		if A[index]%2 == 0 {
			total += A[index]
		}
		res[i] = total
	}
	return res
}