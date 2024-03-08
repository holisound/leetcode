package lc1887

func reductionOperations(nums []int) int {
	cnt := map[int]int{}
	keys := []int{}
	for _, x := range nums {
		if cnt[x]++; cnt[x] == 1 {
			keys = append(keys, x)
		}
	}
	tot := 0
	sort.Ints(keys)
	suffix := 0
	for i := len(keys) - 1; i > 0; i-- {
		suffix += cnt[keys[i]]
		tot += suffix
	}
	return tot
}
