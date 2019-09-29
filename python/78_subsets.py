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
    
        # 迭代
        res=[[]]
        n = len(nums)
        for i in range(n):
            for j in range(len(res)):
                res.append(res[j]+[nums[i]])
        return res
