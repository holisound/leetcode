# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def reverseBetween(self, head, m, n):
        pre=None
        cur=head
        i=0
        front=head
        tail=back=None
        while cur:
            if m<=i<=n:
                tmp=cur.next
                if pre is None:
                    tail=cur
                cur.next=pre
                    
                pre=cur
                cur=tmp
            else:
                if i==m-1:
                    front=cur
                elif i==n+1:
                    back=cur
                cur=cur.next
            i+=1
        front.next=pre
        tail.next=back
        return head
