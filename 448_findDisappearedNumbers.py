class Solution(object):
    def findDisappearedNumbers(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        L = len(nums)
        for i in range(L):
            while nums[i] != nums[nums[i]-1]:
                nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
                print nums
        ans = []
        for i in range(L):
            if i + 1 != nums[i]:
                ans.append(i+1)
        return ans
        

if __name__ == '__main__':
    Solution().findDisappearedNumbers([4,3,2,7,8,2,3,1])