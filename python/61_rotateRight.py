# coding:utf-8
# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

def get_head(arr):
    dummy = ListNode(0)
    cur = dummy
    for x in arr:
        cur.next = ListNode(x)
        cur = cur.next
    return dummy.next

def read_head(head):
    res = []
    node = head
    while node:
        res.append(node.val)
        node= node.next
    return res

class Solution(object):
    def rotateRight(self, head, k):
        if k == 0 or head is None:
            return head
        
        dummy=ListNode(0)
        dummy.next=head
        p=dummy
        length=0
        
        while p.next:                       # 得到list的长度
            length+=1
            p=p.next
        
        p.next=dummy.next
        
        step=length-(k%length)              # 得到新head的位置
        
        while step>0:                       # 到新head的位置
            step-=1
            p=p.next
        
        dummy.next=p.next                   # 变化
        p.next=None 
        
        return dummy.next
    def rotateRightMe(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        if k == 0 or head is None:
            return head

        node = head
        cnt = 0
        tail = None
        while node:
            tail = node
            cnt += 1
            node = node.next
        # -----
        k %= cnt
        j = cnt - k
        dummy = ListNode(0)
        cur = dummy
        node = head

        while j > 0:
            cur.next = ListNode(node.val)
            cur = cur.next
            node = node.next
            j -= 1

        res = node
        while node.next:
            node = node.next
        node.next = dummy.next
        # res = node.next
        return res
if __name__ == '__main__':
    print read_head(Solution().rotateRight(get_head([1,2,3,4,5]),2))
    print read_head(Solution().rotateRight(get_head([1,2]),1))
    print '-'*20
    print read_head(Solution().rotateRightMe(get_head([1,2,3,4,5]),2))
    print read_head(Solution().rotateRightMe(get_head([1,2]),1))