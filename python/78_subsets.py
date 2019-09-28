class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        def backtrack(nums, p, last, res):
            n = len(nums)
            for i in range(p,n):
                temp=last+[nums[i]]
                res.append(temp)
                backtrack(nums, i+1, temp, res)
            return res
        
        return backtrack(nums, 0, [], [[]])
