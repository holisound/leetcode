class Solution(object):
    def shortestPathBinaryMatrix(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        N = len(grid)
        que = collections.deque()
        que.append([0,0,1])
        
        if grid[0][0] == 1: # 判断(0, 0)位置是否为0
            return -1
        while que:
            r, c, p = que.popleft()
            if r == N-1 and c == N-1:
                return p
            for dr in [-1, 0, 1]:
                for dc in [-1, 0, 1]:
                    if dr == 0 and dc == 0:
                        continue
                    if r+dr > -1 and r+dr < N and c+dc > -1 and c+dc < N:
                        if grid[r+dr][c+dc] == 0: # 0 表示可以通行
                            que.append([r+dr, c+dc, p+1])
                            grid[r+dr][c+dc] = 1 # 防止重复遍历
            
        return -1
