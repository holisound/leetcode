class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        res = []
        if not nums:
            return res
        stack = [[nums[0]]]
        j = 0
        for i in range(1, len(nums)):
            if stack[j][-1] + 1 == nums[i]:
                if len(stack[j]) > 1:
                    stack[j].pop()
                stack[j].append(nums[i])
            else:
                j += 1
                stack.append([nums[i]])
        return ['->'.join(str(n) for n in s) for s in stack]
