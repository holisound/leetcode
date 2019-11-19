class Solution(object):
    def findPairs(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        # k = abs(i - j)
        # k >= 0
        # k = i - j
        # k - i = -j
        # i - k = j
        if k < 0:
            return 0
        seen = set()
        pairs = set()
        count = 0
        for i in nums:
            for j in [i-k, i+k]:
                if j in seen and (i, j) not in pairs and (j, i) not in pairs:
                    count += 1
                    pairs.add((i, j))
            seen.add(i)

        return count

if __name__ == '__main__':
    for nums, k in [
    ([3, 1, 4, 1, 5], 2),
    ([1, 2, 3, 4, 5], 1),
    ([1, 2, 3, 4, 5], 3),
    ([1, 3, 1, 5, 4], 0),
    ([6,3,5,7,2,3,3,8,2,4], 2),
    ]:
        print Solution().findPairs(nums, k)