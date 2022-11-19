class Solution(object):
    def permute(self, nums):
        def backtrack(nums, p, res):
            n = len(nums)
            if p == n:
                res.append(list(nums))
            for i in range(p, n):
                nums[i],nums[p]=nums[p],nums[i]
                backtrack(nums, p+1, res)
                nums[i],nums[p]=nums[p],nums[i]
            return res
        return backtrack(nums, 0, [])
    
        # 迭代
        res=[nums]
        n = len(nums)
        for i in range(1,n):
            for j in range(len(res)):
                for k in range(i):
                    temp=list(res[j])
                    temp[k],temp[i]=temp[i],temp[k]
                    res.append(temp)
        return res
