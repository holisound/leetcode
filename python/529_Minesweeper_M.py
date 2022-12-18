
class Solution:
    def updateBoard(self, board: List[List[str]], click: List[int]) -> List[List[str]]:
        R, C = len(board), len(board[0])
        r, c = click

        def neighbors(r, c):
            for dr in [-1, 0, 1]:
                for dc in [-1, 0, 1]:
                    if dr == 0 and dc == 0:
                        continue
                    nr, nc = r + dr, c + dc
                    if -1 < nr < R and -1 < nc < C:
                        yield nr, nc
        if board[r][c] == 'M':
            board[r][c] = 'X'
        else:
            que = collections.deque([(r, c)])
            while que:
                i, j = que.popleft()
                if board[i][j] != 'E':
                    continue  # P.S
                count = 0
                nebs = list(neighbors(i, j))
                for x, y in nebs:
                    if board[x][y] == 'M':
                        count += 1
                board[i][j] = 'B' if count == 0 else str(count)
                if board[i][j] == 'B':  # P.S
                    for x, y in nebs:
                        que.append((x, y))

        return board
