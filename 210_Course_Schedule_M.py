class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        '''参考<python 算法教程> p91 计数入边'''
        G = collections.defaultdict(list)
        count = {num: 0 for num in range(numCourses)}
        for u, v in prerequisites:
            count[u] += 1
            G[v].append(u)
        Q = [u for u in count if count[u] == 0]
        S = []
        while Q:
            u = Q.pop()
            S.append(u)
            for v in G[u]:
                count[v] -= 1
                if count[v] == 0:
                    Q.append(v)
        # 最后判断是否能完成所有课程?
        return S if len(S) == numCourses else []
