class Solution(object):
    def checkPossibility(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        prev = None
        cnt = 0
        for i in range(len(nums)):
            n = nums[i]
            if prev:
                if prev > n:
                    cnt += 1
                    nums[i - 1] = n - 1
                    if i-2 >= 0:
                        if nums[i - 2] > nums[i-1]:
                            return False
            prev = n
        return cnt <= 1

if __name__ == '__main__':
    for s in [[4,2,3], [4,2,1], [3,4,2,3], [1,1,1],[2,3,3,2,4]]:
        print s, Solution().checkPossibility(s)