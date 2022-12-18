class Solution:
    def flipgame(self, fronts: List[int], backs: List[int]) -> int:
        # 记录每张牌，正面和背面数字的出现最大次数
        # 想要的牌的最大出现次数必然是<=1
        n = len(fronts)
        dim = collections.Counter()
        for i in range(n):
            f, b = fronts[i], backs[i]
            num = 2 if f == b else 1
            dim[f] = max(dim[f], num)
            dim[b] = max(dim[b], num)
        res = float('inf')
        for i in range(n):
            if dim[fronts[i]] <= 1:
                res = min(res, fronts[i])
            if dim[backs[i]] <= 1:
                res = min(res, backs[i])
        if res == float('inf'):
            return 0
        return res
