package lc3085

func minimumDeletions(word string, k int) int {
	// 元素集合中任意2个元素频率之差<=k
	cnt := map[rune]int{}
	for _, c:=range word{
		cnt[c]++
	}
	minFreq := len(word)
	for _, freq := range cnt {
		minFreq = min(minFreq, freq)
	}
	ans := 0
	for _, freq := range cnt {
		if diff := freq - (minFreq + k); diff > 0 {
			ans += min(diff, freq)
		}
	}
	return ans

}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
//a:4, b:2, c:1       k = 0
// a:1, b:2, c: 3, d:5  k = 2
// a:6, b:1		k=2