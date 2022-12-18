class Solution(object):
    def removeKdigits(self, num, k):
        """
        :type num: str
        :type k: int
        :rtype: str
        """
        stack=[]
        for c in num:
            while k > 0 and stack:
                if stack[-1] > c:
                    stack.pop()
                    k-=1
                else:
                    break
            if not stack and c == '0': continue
            stack.append(c)
        while k > 0 and stack:
            stack.pop()
            k-=1
        if len(stack)==0:
            return '0'
        return ''.join(stack)
