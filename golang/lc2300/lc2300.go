package lc2300

import "sort"

// https://leetcode.cn/problems/successful-pairs-of-spells-and-potions
func successfulPairs(spells []int, potions []int, success int64) []int {
	n, m := len(spells), len(potions)
	sort.Ints(potions)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		// cond := int((success + int64(spells[i])-1)/int64(spells[i]))
		// (a+b-1)/b = (a-1)/b+1
		res[i] = m - sort.SearchInts(potions, (int(success)-1)/spells[i]+1)
	}

	return res
}
