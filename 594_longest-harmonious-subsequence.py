import collections
import timeit

class Solution(object):
    def findLHS(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        cache = {}
        for n in nums:
            a = n - 1
            b = n + 1
            if n in cache:
                cache[n]['arr'][0] += 1
                cache[n]['arr'][1] += 1
            else:
                d = cache[n] = {}
                d['arr'] = [1, 1]
                d['flag'] = False
            if n+1 in cache:
                cache[n + 1]['arr'][0] += 1
                cache[n + 1]['flag'] = True
            if n-1 in cache:
                cache[n - 1]['arr'][1] += 1
                cache[n - 1]['flag'] = True
        valid = [v['arr'] for v in cache.values() if v['flag']]
        if not valid:
            return 0
        return max(max(v) for v in valid)

    def findLHSv2(self,  nums):
        def findLHS(self, nums):
            counts=collections.Counter(nums)
            return max([counts[i]+counts[i+1] for i in counts if counts[i+1]] or [0])
        # count = collections.Counter(nums)
        # sort_items = sorted(count.items())
        # ans = 0
        # for i in range(1, len(sort_items)):
        #     pre = sort_items[i - 1]
        #     cur = sort_items[i]
        #     if cur[0] - pre[0] == 1:
        #         ans = max(cur[1] + pre[1], ans)
        # return ans

if __name__ == '__main__':
    S = Solution()
    sample = [1,3,2,2,5,2,3,7]
    print timeit.timeit('S.findLHS(sample)', 'from __main__ import S, sample', number=10000)
    print timeit.timeit('S.findLHSv2(sample)', 'from __main__ import S, sample', number=10000)