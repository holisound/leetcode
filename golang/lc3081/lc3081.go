package lc3081

func minimizeStringValue(s string) string {
	cnt := make([]int, 26)
	cntQ := make([]int, 26)
	chars := []rune(s)

	for _, c := range s {
		if c != '?' {
			cnt[c-'a']++
		}
	}
	for _, c := range s {
		if c != '?' {
			continue
		}
		// 枚举26字母，取频次最小的字母
		freqMin := len(s)
		cond := 0
		for i := 0; i < 26; i++ {
			if cnt[i] < freqMin {
				freqMin = cnt[i]
				cond = i
			}
		}
		cnt[cond]++
		cntQ[cond]++
	}

	for j, c := range chars {
		if c != '?' {
			continue
		}
		// 按字典序替换"?"
		for i := 0; i < 26; i++ {
			if cntQ[i] > 0 {
				cntQ[i]--
				chars[j] = rune(int('a') + i)
				break
			}
		}
	}
	return string(chars)
}
