func commonChars(A []string) []string {
	/*
		这题我学到了:
		合理的使用packages, 这里使用strings.Count
	*/
	res := []string{}
	seen := map[rune]bool{}
	for _, k := range A[0] {
		if seen[k] {
			continue
		}
		seen[k] = true
		n := math.MaxInt32
		for _, w := range A {
			n = min(n, strings.Count(w, string(k)))
		}
		for i := 0; i < n; i++ {
			res = append(res, string(k))
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}