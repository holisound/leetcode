class Solution:
    def restoreIpAddresses(self, s: str) -> List[str]:
        def backtrack(i, dot, last, res):
            if i == len(s):
                if dot == 3 and ok(last):
                    res.append(last)
                return
            if dot < 3:
                backtrack(i + 1, dot + 1, last + s[i] + '.', res)
            backtrack(i + 1, dot, last + s[i], res)

            return res
        if len(s) > 12:
            return []
        return backtrack(0, 0, '', [])


def ok(s):
    for x in s.split('.'):
        if not x:
            return False
        if len(x) > 1 and x[0] == '0':
            return False
        if int(x) < 0 or int(x) > 255:
            return False
    return True
