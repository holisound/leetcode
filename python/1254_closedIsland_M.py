class Solution:
    def closedIsland(self, grid: List[List[int]]) -> int:
        m = len(grid)
        n = len(grid[0])
        visited = [[0] * n for _ in range(m)]

        def dfs(i, j):
            if (i == 0 or i == m - 1 or j == 0 or j == n - 1) and grid[i][j] == 0:
                return 0
            res = 1
            if not visited[i][j]:
                visited[i][j] = 1
                if grid[i][j] == 0:
                    # grid[i][j]=1
                    for r, c in [(i + 1, j), (i, j + 1), (i - 1, j), (i, j - 1)]:
                        if -1 < r < m and -1 < c < n:
                            res &= dfs(r, c)
            return res
        res = 0
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 0 and visited[i][j] == 0:
                    if dfs(i, j):
                        res += 1
        return res
