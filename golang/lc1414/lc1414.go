package lc1414

func findMinFibonacciNumbers(k int) int {
	fibs := []int{}
	for a, b := 0, 1; b <= k; a, b = b, a+b {
		fibs = append(fibs, b)
	}
	ans := 0
	for i := len(fibs) - 1; i > -1 && k != 0; i-- {
		if fibs[i] <= k {
			k -= fibs[i]
			ans++
		}
	}
	return ans
}
