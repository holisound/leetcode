class Solution:
    def findRepeatedDnaSequences(self, s: str) -> List[str]:
        d={}
        for i in range(len(s)-10+1):
            win = s[i:i+10]
            if win in d:
                d[win]=True
            else:
                d[win]=False
        return [k for k,v in d.items() if v]
