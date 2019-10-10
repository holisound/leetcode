class Solution:
    def readBinaryWatch(self, num: int) -> List[str]:
        # 回溯法
        # 相当求 count(1)==num,所有二进制数（位数<= 10)
        def backtrack(digs, count, last, res):
            if digs > 10:
                if count == num:
                    h = last >> 6
                    m = last - (h << 6)
                    if 0 <= h < 12 and 0 <= m < 60:
                        res.append("%d:%02d" % (h, m))
            else:
                last <<= 1
                backtrack(digs + 1, count + 1, last | 1, res)
                backtrack(digs + 1, count, last, res)
            return res
        return backtrack(0, 0, 0, [])

        # 迭代版本
        que=collections.deque()
        que.append((0,0,0))
        res = []
        while que:
            last, cnt, depth = que.popleft()
            if depth > 10:
                if cnt==num:
                    h = last >> 6
                    m = last - (h << 6)
                    if 0 <= h < 12 and 0 <= m < 60:
                        res.append("%d:%02d" % (h, m))
            else:
                last <<= 1
                que.append((last|1, cnt+1, depth+1))
                que.append((last, cnt, depth+1))

        return res
