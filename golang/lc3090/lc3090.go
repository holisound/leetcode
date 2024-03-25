package lc3090

func maximumLengthSubstring(s string) int {
	left, right := [26]int{}, [26]int{}
	ans := 0
	j := 0
	for i, c := range s {
		if left[c-'a'] == 0 {
			left[c-'a'] = i + 1
		} else if right[c-'a'] == 0 {
			right[c-'a'] = i + 1
		} else {
			j = max(j, left[c-'a'])
			left[c-'a'], right[c-'a'] = right[c-'a'], i+1
		}
		ans = max(ans, i-j+1)
	}
	return ans
}
