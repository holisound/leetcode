class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        N = len(nums)
        F = [-1 << 31] * (N + 1)
        maxhere = -1 << 31
        for i in range(1, N + 1):
            maxhere = max(nums[i - 1], maxhere + nums[i - 1])
            F[i] = max(F[i - 1], maxhere)
        return F[N]
