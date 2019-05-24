/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }

    count, cur := 0, head
    for cur != nil {
        count++
        cur = cur.Next
    }
    k %= count
    if k == 0 {
        return head
    }

    phead, tail := head, head
    for i := 0; i < count-k; i++ {
        tail = phead
        phead = phead.Next
    }
    tail.Next = nil
    dm := &ListNode{0, nil}
    cur = dm
    for phead != nil {
        cur.Next = phead
        cur = cur.Next
        phead = phead.Next
    }
    cur.Next = head
    return dm.Next
}