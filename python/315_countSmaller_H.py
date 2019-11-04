class Solution:
    def countSmaller(self, nums: List[int]) -> List[int]:
        # O(N^2) TLE
        count=collections.Counter()
        for i,n in enumerate(nums):
            for index in count:
                if n < nums[index]:
                    count[index]+=1
            count[i] = 0
        return [count[i] for i in range(len(nums))]

        # O(NLogN) AC, inspired by LIS problem
        lis = []
        res = []

        for num in nums[::-1]:
            idx = bisect.bisect_left(lis, num)
            res.append(idx)
            lis.insert(idx, num)
        return res[::-1]
