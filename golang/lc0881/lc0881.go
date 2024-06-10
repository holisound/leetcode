package main

import "sort"

func numRescueBoats(people []int, limit int) int {
	ans := len(people)
	n := len(people)
	sort.Ints(people)
	index := map[int]bool{}

	for i, p := range people {
		if index[i] {
			continue
		}

		idx := sort.Search(n-1, func(mid int) bool {
			return people[mid+1] > limit-p
		})

		// if idx > i && people[idx]+p <= limit {
		for j := idx; j >= i+1; j-- {
			if !index[j] {
				index[i] = true
				index[j] = true
				ans--
				break
			}
		}
		// }
	}

	return ans
}
