class Solution(object):
    def allPathsSourceTarget(self, graph):
        N=len(graph)
        # 自顶向下
        def upDown(graph,last,result):
            if last[-1]==N-1:
                result.append(last)
                return
            u=last[-1]
            for v in graph[u]:
                upDown(graph, last+[v], result)
            return result
        return upDown(graph,[0],[])
        # 自底向上
        def solve(node):
            if node == N-1: return [[N-1]]
            ans = []
            for nei in graph[node]:
                for path in solve(nei):
                    ans.append([node] + path)
            return ans
        
