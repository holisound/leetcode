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

def count_dyck_word(n, x, y):
    count_dyck_word.count = 0
    count_dyck_word.enum = []

    def backtrack(l, r, p):
        if r > l or l > n:
            return
        if l == r == n:
            count_dyck_word.enum.append(p)
            count_dyck_word.count += 1
        backtrack(l + 1, r, p + x)
        backtrack(l, r + 1, p + y)
    backtrack(0, 0, '')
    return count_dyck_word
