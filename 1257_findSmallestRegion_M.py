class Solution:
    def findSmallestRegion(self, regions: List[List[str]], region1: str, region2: str) -> str:
        # solution #1 Graph, DFS (up-down)
        G = collections.defaultdict(set)
        for regs in regions:
            u, v = regs[0], regs[1:]
            G[u].update(set(v))

        def depth(u, t, d):
            if u == t:
                return d + 1
            if u in G:
                return max(depth(v, t, d + 1) for v in G[u])
            return -1
        d, res = float('inf'), ''
        for u in G:
            d1 = depth(u, region1, -1)
            d2 = depth(u, region2, -1)
            # print(d1,d2,u)
            if (d1 != -1 and d2 != -1) and d1 + d2 < d:
                d, res = d1 + d2, u
        return res
        # solution #2 hashmap, down-up
        dic = {}
        for regs in regions:
            for reg in regs[1:]:
                dic[reg] = regs[0]

        s1, s2 = set(), set()
        while 1:
            r1, r2 = region1, region2
            s1.add(r1)
            s2.add(r2)
            if r1 in s2:
                return r1
            if r2 in s1:
                return r2
            region1, region2 = dic.get(r1), dic.get(r2)
