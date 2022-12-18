class Solution(object):
    def findUnsortedSubarray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        L = len(nums)
        sn = sorted(nums)
        if nums == sn:
            return 0
        maxi = L - 1
        a, b = None, None
        for i in range(L):
            j = maxi - i
            if j == 0:
                break
            if a is None:
                if nums[i] != sn[i]:
                    a = i
            if b is None:
                if nums[j] != sn[j]:
                    b = j
            if a and b:
                break
        return b - a + 1
            
if __name__ == '__main__':
    Solution().findUnsortedSubarray([1,2,3,4])