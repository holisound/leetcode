package lc1625

func Add(s string, a, start int) string {
	chars := []byte(s)
	for i := start; i < len(s); i += 2 {
		k := int(chars[i] - '0')
		chars[i] = byte('0' + (k+a)%10)
	}
	return string(chars)
}

func findLexSmallestString(s string, a int, b int) string {
	/*
	s[i] 累加 "a" 之后的数位 只可能是0-9之一
	s[i]（一个数位）累加"a"10次,将回到s[i] (因为(s[i] + a*10) % 10 = s[i])
	*/
	n := len(s)
	ans := s
	for _ = range s {
		s = s[n-b:] + s[:n-b]
		for j := 0; j < 10; j++ {
			s = Add(s, a, 1)
            ans = min(ans, s)
			if b&1 == 1 {
				for p := 0; p < 10; p++ {
					s = Add(s, a, 0)
                    ans = min(ans, s)
				}
			}
		}
	}
	return ans
}
