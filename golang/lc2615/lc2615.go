package main

func distance(nums []int) []int64 {
	n := len(nums)
	ans := make([]int64, n)
	cnt := map[int][]int{}
	for i, x := range nums {
		cnt[x] = append(cnt[x], i)
	}
	for _, arr := range cnt {
		tot := 0
		for i := range arr {
			tot += arr[i]
		}
		m := len(arr)
		l, r := 0, tot
		for i, pos := range arr {
			ans[pos] = int64(i*pos - l + r - (m-i)*pos)
			l, r = l+pos, r-pos
		}
	}

	return ans
}
