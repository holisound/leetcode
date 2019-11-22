# coding:utf-8
class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        # 借鉴两数之和的思路， 使用哈希表
        count = collections.defaultdict(list)

        for j in range(1, len(nums)):
            for i in range(j):
                count[nums[i] + nums[j]].append([i, j])
        res = []
        seen = set()
        for j in range(len(nums)):
            for i in range(j):
                s = target - (nums[i] + nums[j])
                for m, n in count[s]:
                    if i == m or i == n or j == m or j == n:
                        continue
                    con = sorted(nums[k] for k in [m, n, i, j])
                    if tuple(con) not in seen:
                        seen.add(tuple(con))
                        res.append(con)
        return res
