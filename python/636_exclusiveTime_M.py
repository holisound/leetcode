# 2
# ["0:start:0","1:start:2","1:end:5","0:end:6"]
# 1
# ["0:start:0","0:start:2","0:end:5","0:start:6","0:end:6","0:end:7"]
# 2
# ["0:start:0","0:start:2","0:end:5","1:start:6","1:end:6","0:end:7"]


class Solution(object):
    def exclusiveTime(self, n, logs):
        """
        :type n: int
        :type logs: List[str]
        :rtype: List[int]
        """
        stack = []
        res = [0] * n
        for log in logs:
            f, s, t = log.split(':')
            f, t = int(f), int(t)
            if s == "end":
                p = stack.pop()
                v = t - p[1] + 1
                res[f] += v
                if stack:
                    # 减去当前函数的执行时间，使用栈顶id的进行修改, 即外层函数的执行时间
                    res[stack[-1][0]] -= v
            else:
                stack.append((f, t))
        return res
