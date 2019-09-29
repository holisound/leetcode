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
    
        #使用queue迭代
        que=collections.deque([("", 0, 0)])
        res=[]
        while que:
            s,l,r=que.popleft()
            if l > n or r > l:
                continue
            if l==r==n:
                res.append(s)
            que.append((s+"(",l+1,r))
            que.append((s+")",l,r+1))
            
        return res
