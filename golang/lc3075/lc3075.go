package lc3075

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)
	ans := int64(0)
	n := len(happiness)
	for i := 0; i < n && k > 0; i++ {
		h := max(0, happiness[n-i-1] - i)
		ans += int64(h)
		k--
	}
	return ans
}
