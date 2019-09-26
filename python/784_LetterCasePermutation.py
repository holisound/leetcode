class Solution:
    def letterCasePermutation(self, S: str) -> List[str]:
        
        def backtrack(S, start, res):
            n = len(S)
            for i in range(start,n):
                if S[i].isalpha():
                    nc=chr(ord(S[i])^32)
                    ns=S[:i]+nc+S[i+1:]
                    # 注意:下次起始位置应该是i+1, 不是start+1
                    res.extend(backtrack(ns, i+1, [ns])) 

            return res
        return backtrack(S, 0, [S])
