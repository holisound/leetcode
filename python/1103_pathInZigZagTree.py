class Solution(object):
    def pathInZigZagTree(self, label):
        """
        :type label: int
        :rtype: List[int]
        """
        # m = []
        # i = j = 1 # i 表示当前label的值, j 表示当前行节点的个数
        # c = -1
        # while i <= label:
        #     cur = []
        #     for jj in range(j):
        #         cur.append(i)
        #         i += 1
        #         if i > label:
        #             c = jj
        #             break
        #     m.append(cur)
        #     j *= 2
        # r = len(m) - 1
        R = 0
        n = label
        while n > 0:
            R += 1
            n >>= 1
        res = []
        r = R - 1
        c = label - 2**r
        # 自底向上查找
        # print(r,c, m[r][c])
        while r > -1:
            res.append(getLabel(r, c))
            # 这行代码值得细细品味,
            # c是label的列, 根据当前列求出上面一行的列的位置
            c = (2 ** r - c - 1) >> 1
            r -= 1
        return res[::-1]


def getLabel(r, c):
    if r <= 0:
        return 1
    return 2 ** r + c
