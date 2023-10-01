package main

func predictPartyVictory(senate string) string {
	for {
		q1 := []rune{}
		q2 := []rune{}
		for _, c := range senate {
			if len(q1) > 0 && q1[0] != c {
				q2 = append(q2, q1[0])
				q1 = q1[1:]
				continue
			}
			q1 = append(q1, c)
		}
		if len(q2) == 0 {
			break
		}
		senate = string(q1) + string(q2)
	}
	return map[string]string{
		"R": "Radiant", "D": "Dire",
	}[senate[:1]]
}
