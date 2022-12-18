package main

/*
https://leetcode.cn/problems/design-memory-allocator/
2502. 设计内存分配器
*/
type Allocator struct {
	blocks []int
}

func Constructor(n int) Allocator {
	return Allocator{make([]int, n)}
}

func (this *Allocator) Allocate(size int, mID int) int {
	cnt := 0
	start := -1
	for i, b := range this.blocks {
		if b == 0 {
			cnt++
		} else {
			cnt = 0
		}
		if cnt == size {
			start = i - cnt + 1
			break
		}
	}
	if start == -1 {
		return -1
	}
	for i := start; i < start+size; i++ {
		if this.blocks[i] == 0 {
			this.blocks[i] = mID
		}
	}
	return start

}

func (this *Allocator) Free(mID int) int {
	cnt := 0
	for i, b := range this.blocks {
		if b == mID {
			cnt++
			this.blocks[i] = 0
		}
	}
	return cnt
}
