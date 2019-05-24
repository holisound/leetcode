class Solution(object):
    def isSubsequence(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        def bisect_left(a, x):
            lo = 0
            hi = len(a)
            while lo < hi:
                mid = (lo +_hi) // 2
                if a[mid] < x:
                    lo = mid + 1
                elif a[mid] == x:
                    return mid
                else:
                    hi = mid
            return -1

        L = len(t)
        prev = 0
        for c in s:
            pos = bisect_left(t, c)
            if pos == L or pos < prev:
                return False
            prev = pos
        return True

    def greedy(self, s, t)
        if len(s) == 0:
            return True
        if len(t) == 0:
            return False
        index = 0
        for char in t:
            if s[index] == char:
                index += 1
            if index == len(s):
                return True
        return False