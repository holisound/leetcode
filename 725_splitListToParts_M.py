class Solution:
    def splitListToParts(self, root: ListNode, k: int) -> List[ListNode]:
        n, node = 0, root
        while node:
            node, n = node.next, n + 1
        p = [n // k for _ in range(k)]
        for i in range(n % k):
            p[i] += 1
        heads = [ListNode(0) for _ in range(k)]
        now = list(heads)
        while root:
            for i in range(k):
                pre = now[i]
                cur = pre.next = root
                for _ in range(p[i]):
                    pre, cur = cur, cur.next
                pre.next, root = None, cur
        return [h.next for h in heads]
