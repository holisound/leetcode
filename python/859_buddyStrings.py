class Solution(object):
    def buddyStrings(self, A, B):
        """
        :type A: str
        :type B: str
        :rtype: bool
        """
        L = len(A)
        if L == 0 or L != len(B):
            return False
        n = 0
        diff = None
        r = True
        if L > 2:
            r = A == B
        counter = {}
        duplicated = False
        while n < L:
            a = A[n]
            b = B[n]
            if not duplicated and a not in counter:
                counter[a] = 1
            else:
                duplicated = True
                
            if n > 0 and L <= 2:
                r &= a == A[n - 1]
            if a != b:
                if diff is None:
                    diff = a,b
                elif (b, a) == diff:
                    r = True
                else:
                    r = False
            n += 1
        if A == B:
            r &= duplicated
        return r