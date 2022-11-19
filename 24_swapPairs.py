# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def swapPairs(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if not head:
            return head
        a = head
        b = a.next
        last = None
        root = b if b else a
        while b:
            c = b.next
            b.next = a
            a.next = c
            if last:
                last.next = b
            last = a
            a = c
            if not a or not a.next:
                break
            b = a.next
        return root