package lc2780

func minimumIndex(nums []int) int {
	n := len(nums)
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	left := map[int]int{}
	for i, x := range nums {
		left[x]++
		if right := cnt[x] - left[x]; right*2 > (n-i-1) && left[x]*2 > i+1 {
			return i
		}
	}
	return -1
}
