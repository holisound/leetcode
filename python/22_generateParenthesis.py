class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        def backtrack(res, l, r, s):
            if l > n or r > l:
                return
            if l == r == n:
                res.append(s)
            backtrack(res, l + 1, r, s + '(')
            backtrack(res, l, r + 1, s + ')')
            return res
        return backtrack([], 0, 0, "")
