class Solution:
    def corpFlightBookings(self, bookings, n):
        # 为了帮助理解，在二维坐标轴里，先把线段画出来
        # 线段重叠部分数值累加，
        # 到达线段出发点start，加上数，
        # 到达线段终点end，同样加上数，到达end+1,减去数

        count = collections.defaultdict(int)
        for s, e, num in bookings:
            count[s] += num
            count[e] + num
            count[e + 1] -= num

        res = []

        for i in range(1, n + 1):
            if not res:
                res.append(count[i])
            else:
                res.append(count[i] + res[-1])
        return res
