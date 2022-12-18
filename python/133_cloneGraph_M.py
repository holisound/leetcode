"""
# Definition for a Node.
class Node:
    def __init__(self, val, neighbors):
        self.val = val
        self.neighbors = neighbors
"""

class Solution:
    def cloneGraph(self, node: 'Node') -> 'Node':
        cache = {}
        def clone(node):
            if node in cache:
                return cache[node]
            cloned = Node(node.val, [])
            # 我已经意识到需要设置缓存, 但卡在这里 
            # 缓存设置需在递归开始前, 这样才能利用到之前的缓存
            cache[node] = cloned 
            for ne in node.neighbors:
                cloned.neighbors.append(clone(ne))
            return cloned
        return clone(node)
