# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        values = []
        while l1:
            values.append(l1.val)
            l1 = l1.next
        for i in range(len(values)):
            if not l2:
                break
            values[i] += l2.val
            l2 = l2.next
        while l2:
            values.append(l2.val)
            l2 = l2.next
        carry = 0
        for i in range(len(values)):
            s = values[i] + carry
            carry, values[i] = s // 10, s % 10
        if carry > 0:
            values.append(carry)
        dummy = now = ListNode(0)
        for v in values:
            now.next = ListNode(v)
            now = now.next
        return dummy.next
