package lc2947

func beautifulSubstrings(s string, k int) int {
	n := len(s)
	ans := 0
	// 枚举子串
	for i := 1; i <= n; i++ {
		vowels, cons := 0, 0
		for j := 0; j < i; j++ {
			if isVowel(s[j]) {
				vowels++
			} else {
				cons++
			}

			if vowels == cons && (vowels*cons)%k == 0 {
				ans++
			}
		}
	}

	return ans

}

func isVowel(b byte) bool {
	return strings.IndexByte("aeiou", b) != -1
}
