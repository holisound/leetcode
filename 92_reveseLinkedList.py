# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def reverseBetween(self, head: ListNode, m: int, n: int) -> ListNode:
        # The most-left node of the list is unstable
        # (can't make sure the left border)
        # need a dummy head to assist us determine the left border
        dummy = ListNode(0)
        dummy.next = head
        cur, pre = head, dummy
        for _ in range(m - 1):
            pre, cur = cur, cur.next
        rHead, left = cur, pre  # find the first node to reverse
        pre.next = None  # break the chain of the list

        for _ in range(n - m + 1):
            pre, cur = cur, cur.next
        # find the first node of right part
        right = cur
        # break the chain again
        pre.next = None

        middle = reverse(rHead)
        left.next, rHead.next = middle, right  # rHead now is at the end
        return dummy.next


def reverse(head):
    pre = None
    while head:
        temp = head.next
        head.next = pre
        pre = head
        head = temp
    return pre
