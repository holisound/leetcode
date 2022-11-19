class Solution:
    def letterCasePermutation(self, S: str) -> List[str]:
        
        def backtrack(S, start, res):
            n = len(S)
            for i in range(start,n):
                if S[i].isalpha():
                    nc=chr(ord(S[i])^32)
                    ns=S[:i]+nc+S[i+1:]
                    # 注意:下次起始位置应该是i+1, 不是start+1
                    res.append(ns)
                    backtrack(ns, i+1, res)

            return res
        return backtrack(S, 0, [S])
    
        def backtrack2(p, res, n):
            k=-1
            for i in range(len(res)):
                x=res[i]
                for j in range(p, n):
                    if x[j].isalpha():
                        nx=x[:j]+chr(ord(x[j])^32)+x[j+1:]
                        res.append(nx)
                        k=j+1
                        break
            if k !=-1:
                backtrack2(k, res, n)
            return res
            
        return backtrack2(0,[S], len(S))
        
        def backtrack3(chars, p, res):
            n = len(chars)
            for i in range(p, n):
                if chars[i].isalpha():
                    temp = chars[i]
                    chars[i] = chr(ord(chars[i]) ^ 32)
                    res.append(''.join(chars))
                    backtrack3(chars, i+1, res)
                    chars[i] = temp
            return res
        return backtrack3(list(S), 0, [S])
        # 迭代版本
        res=[S]
        n = len(S)
        for i in range(n):
            for j in range(len(res)):
                x=res[j]
                if x[i].isalpha():
                    res.append(x[:i]+chr(ord(x[i])^32)+x[i+1:])
        return res
