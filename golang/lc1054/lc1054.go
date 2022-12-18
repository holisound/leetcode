package main

import "sort"

func rearrangeBarcodes(barcodes []int) []int {
	/*
		这题我学到了：
		很实用的一个例子，如何让相邻的节点不重复
	*/
	count := map[int]int{}
	keys := []int{}
	for _, b := range barcodes {
		count[b]++
		if count[b] == 1 {
			keys = append(keys, b)
		}
	}
	sort.Slice(keys, func(i, j int) bool {
		return count[keys[i]] > count[keys[j]]
	})
	res := make([]int, len(barcodes))
	start, index := 0, 0
	for index < len(keys) {
		k := keys[index]
		for i := start; i < len(res); i += 2 {
			res[i] = k
			count[k]--
			if count[k] == 0 {
				index++
				if index >= len(keys) {
					break
				}
				k = keys[index]
			}
		}
		start++

	}
	return res
}
