package lc2870

func minOperations(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	tot := 0
	for _, v := range cnt {
		if v == 1 {
			return -1
		}
		tot += v / 3
		if v%3 != 0 {
			tot++
		}
	}
	return tot
}
