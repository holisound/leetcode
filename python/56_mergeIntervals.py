class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        res = []
        i = 0
        intervals.sort()
        n = len(intervals)
        while i < n:
            start, end = intervals[i]
            for j in range(i + 1, n):
                s, e = intervals[j]
                if s <= end:  # 区间有交集
                    end = max(end, e)  # 踩坑，取较大值
                    i += 1
                else:
                    break
            res.append([start, end])
            i += 1
        return res
