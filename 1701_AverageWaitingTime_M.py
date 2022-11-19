class Solution:
    def averageWaitingTime(self, customers: List[List[int]]) -> float:
        # 计算实际给每个顾客做菜所花费的时间总和
        # 需要额外处理的情况即是: arr_time(i) + duration(i) > arr_time(i+1)
        # $cur_time 表示当前时间
        total = cur_time = 0
        for arr_time, duration in customers:
            cur_time = max(cur_time, arr_time) + duration  # 仔细体会下，这行代码很妙
            total += cur_time - arr_time
        return total / len(customers)
