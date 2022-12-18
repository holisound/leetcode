package main

import (
	. "github.com/holisound/leetcode/utils"
)

func addBinary(a string, b string) string {

	var carry byte = '0'

	bits := []byte{}

	for i, j := len(a)-1, len(b)-1; carry > '0' ||
		i > -1 || j > -1; i, j = i-1, j-1 {
		count := map[byte]int{}
		count[carry]++
		if i > -1 {
			count[a[i]]++
		}
		if j > -1 {
			count[b[j]]++
		}
		if count['1'] == 3 {
			carry, bits = '1', append(bits, '1')
		} else if count['1'] == 2 {
			carry, bits = '1', append(bits, '0')
		} else if count['1'] == 1 {
			carry, bits = '0', append(bits, '1')
		} else {
			carry, bits = '0', append(bits, '0')
		}
	}
	Reverse(bits)
	return string(bits)
}
