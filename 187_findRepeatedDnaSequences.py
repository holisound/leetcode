class Solution:
    def findRepeatedDnaSequences(self, s: str) -> List[str]:
        res, seen=set(),set()
        for i in range(len(s)-10+1):
            win = s[i:i+10]
            if win in seen:
                res.add(win)
            seen.add(win)
        return list(res)
