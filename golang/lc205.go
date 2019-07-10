func isIsomorphic(s string, t string) bool {
	s2t, t2s := map[byte]byte{}, map[byte]byte{}

	for i := 0; i < len(s); i++ {
		si, ti := s[i], t[i]
		if _, ok := s2t[si]; ok && s2t[si] != ti {
			return false
		}
		if _, ok := t2s[ti]; ok && t2s[ti] != si {
			return false
		}
		s2t[si], t2s[ti] = ti, si
	}
	return true
}