class Solution:
    def circularArrayLoop(self, nums: List[int]) -> bool:
        G = {}
        n = len(nums)
        for i, num in enumerate(nums):
            G[i] = num + i
        for i in range(n):
            visited = loop(G, i, n)
            if len(visited) > 2 and visited[0] == visited[-1]:
                return True
        return False


def loop(G, u, n):
    visited, seen = [], set()
    now = last = 0
    while u not in seen:
        seen.add(u)
        now, last = now + G[u] - u, now
        if abs(now) <= abs(last):
            break
        visited.append(u % n)
        u = G[u] % n
    visited.append(u % n)
    return visited
