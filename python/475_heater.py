class Solution(object):
    def findRadius(self, houses, heaters):
        """
        :type houses: List[int]
        :type heaters: List[int]
        :rtype: int
        """
        if len(heaters) == 0 or len(houses) == 0:
            return
        if len(heaters) < 2:
            return max(houses[-1] - heaters[0], heaters[0] - 1)
        a = self.bisect_left(houses, heaters[0])
        b = self.bisect_left(houses, heaters[-1])
        tail = houses[-1] - heaters[-1]
        r = max(a, tail)
        for h in heaters:
            b = self.bisect_left(houses, h)
            r = max((b-a)/2, r)
            a = b
        return r 

    def bisect_left(self, a, x):
        lo = 0
        hi = len(a)
        while lo < hi:
            mid = (lo + hi) // 2
            if a[mid] < x:
                lo = mid + 1
            else:
                hi = mid
        return lo
    
if __name__ == '__main__':
    for houses, heaters, exp in [
    ([1,2,3], [2], 1),
    ([1,2,3,4], [1,4], 1),
    ([1,5],[10],9),
    ([1,5],[2],3),
    ([1,2,3,5,15],[2,30], 13)
    ]:
        print 'REAL:%s' % Solution().findRadius(houses, heaters), 'EXPECT:%s' % exp