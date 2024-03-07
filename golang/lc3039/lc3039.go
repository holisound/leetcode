package lc3039

func lastNonEmptyString(s string) string {
	cnt := map[rune]int{}
	freqMost := 0
	for _, c := range s {
		cnt[c]++
		freqMost = max(freqMost, cnt[c])
	}
	strb := new(strings.Builder)
	cnt = map[rune]int{}
	for _, c := range s {
		cnt[c]++
		if cnt[c] == freqMost {
			strb.WriteRune(c)
		}
	}

	return strb.String()
}
