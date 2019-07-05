class Solution(object):
    def minHeightShelves(self, books, shelf_width):
        # 动态规划问题的关键在于，你能不能真正地理解问题？
        # w 累加，w <= shelf_width, 可以摆放在同一层
        # 只考虑书架每一层摆放情况有哪些，
        # this_level_books_width <= shelf_width
        # books[i][0] + books[i+1][0] ... + books[i+n][0] <= shelf_width 
        n = len(books)
        ans = [float('inf')]*(n+1)
        ans[0] = 0
        
        for i in range(1, n+1): # dp[i], 表示摆放第i本书的书架高度
            w = 0
            h = 0
            for j in range(i-1, n): # 
                w += books[j][0]
                if w > shelf_width:
                    break
                h = max(h, books[j][1])
                ans[i] = min(ans[i], ans[j]+h)
        return ans[n]

from functools import lru_cache
class Solution:
    def minHeightShelves(self, books, shelf_width: int):
        @lru_cache(None)
        def dp(b, e):
            if b == e: return 0
            res,w,h = float('inf'),0,0
            for i in range(b,e):
                if w + books[i][0] <= shelf_width:
                    h = max(h, books[i][1])
                    w += books[i][0]
                    res = min(res, h + dp(i+1, e))
                    #print(res)
                else:
                    break
            return res
        return dp(0, len(books))
