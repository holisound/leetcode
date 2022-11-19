class Solution:
    def findReplaceString(self, S: str, indexes: List[int], sources: List[str], targets: List[str]) -> str:
        pos = {index: i for i, index in enumerate(indexes)}
        indexes.sort()
        last = 0
        m = {}
        for index in indexes + [len(S)]:
            m[last] = S[last:index]
            last = index

        for index in indexes:
            i = pos[index]
            temp = m[index]
            s, t = sources[i], targets[i]
            if temp.startswith(s):
                m[index] = t + temp[len(s):]
        res = ''
        for k in sorted(m):
            res += m[k]
        return res
