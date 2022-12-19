package main

import (
	"fmt"

	. "github.com/holisound/leetcode/utils"
)

/*
https://leetcode.cn/problems/design-front-middle-back-queue/
1670. 设计前中后队列
*/
type FrontMiddleBackQueue struct {
	Dummy *ListNode
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{&ListNode{}}
}
func (this *FrontMiddleBackQueue) log() {
	cur := this.Dummy.Next
	fmt.Println("--------------")
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
func (this *FrontMiddleBackQueue) PushFront(val int) {
	head := &ListNode{Val: val}
	this.Dummy.Next, head.Next = head, this.Dummy.Next
	// this.log()
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	slow, fast := this.Dummy, this.Dummy.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	tmp := slow.Next
	slow.Next = &ListNode{Val: val}
	slow.Next.Next = tmp
	// this.log()
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	cur := this.Dummy
	for cur != nil && cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{Val: val}
	// this.log()
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.Dummy.Next == nil {
		return -1
	}
	node := this.Dummy.Next
	this.Dummy.Next = node.Next
	return node.Val

}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.Dummy.Next == nil {
		return -1
	}
	pre, slow, fast := this.Dummy, this.Dummy, this.Dummy
	for fast != nil && fast.Next != nil {
		pre, slow, fast = slow, slow.Next, fast.Next.Next
	}
	pre.Next = slow.Next
	return slow.Val
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.Dummy.Next == nil {
		return -1
	}
	pre, cur := this.Dummy, this.Dummy.Next
	for cur != nil && cur.Next != nil {
		pre, cur = cur, cur.Next
	}
	pre.Next = nil
	return cur.Val
}
