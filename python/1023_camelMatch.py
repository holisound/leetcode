class Solution(object):
    def camelMatch(self, queries, pattern):
        """
        :type queries: List[str]
        :type pattern: str
        :rtype: List[bool]
        """
        res = []
        countp = countUpper(pattern)
        for q in queries:
            if countUpper(q) == countp and isSub(pattern, q):
                res.append(True)
            else:
                res.append(False)
        return res


def countUpper(s):
    n = 0
    for c in s:
        if 'A' <= c <= 'Z':
            n += 1
    return n


def isSub(a, b):

    i = 0
    for j in range(len(b)):
        if i < len(a) and b[j] == a[i]:
            i += 1

    return len(a) == i
