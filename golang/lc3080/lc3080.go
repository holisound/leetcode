package lc3080

type element struct {
	Value int
	Index int
}

func unmarkedSumArray(nums []int, queries [][]int) []int64 {
	tot := 0
	n := len(nums)
	elems := []element{}

	for i, x := range nums {
		tot += x
		elems = append(elems, element{x, i})
	}
	sort.Slice(elems, func(i, j int) bool {
		if elems[i].Value == elems[j].Value {
			return elems[i].Index < elems[j].Index
		}
		return elems[i].Value < elems[j].Value
	})
	ans := make([]int64, len(queries))
	j := 0
	sub := 0
	m := make([]int, n)
	for i, q := range queries {
		if m[q[0]] == 0 {
			m[q[0]] = 1
			sub += nums[q[0]]
		}
		for k := q[1]; j < n && k != 0; j++ {
			if idx := elems[j].Index; m[idx] == 0 {
				sub += elems[j].Value
				m[idx] = 1
				k--
			}

		}
		ans[i] = int64(tot - sub)
	}

	return ans
}
