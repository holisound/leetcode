class Solution:
    def findRightInterval(self, intervals: List[List[int]]) -> List[int]:
        n = len(intervals)
        res = [-1] * n
        index = sorted(range(n), key=lambda i: intervals[i])
        intervals.sort()

        for i, it in enumerate(intervals):
            lo, hi = 0, n - 1
            while lo < hi:
                mid = (lo + hi) // 2
                a = intervals[mid][0]
                if a < it[1]:
                    lo = mid + 1
                else:
                    hi = mid
            if intervals[lo][0] >= it[1]:
                res[index[i]] = index[lo]

        return res
