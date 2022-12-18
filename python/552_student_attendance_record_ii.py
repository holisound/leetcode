# coding:utf-8
# 学生出勤记录是只包含以下三个字符的字符串：

# 'A' : Absent，缺勤
# 'L' : Late，迟到
# 'P' : Present，到场
# 如果记录不包含多于一个'A'（缺勤）或超过两个连续的'L'（迟到），则该记录被视为可奖励的。
# 输入: n = 2
# 输出: 8 
# 解释：
# 有8个长度为2的记录将被视为可奖励：
# "PP" , "AP", "PA", "LP", "PL", "AL", "LA", "LL"
# 只有"AA"不会被视为可奖励，因为缺勤次数超过一次。
# n = 3, (27 - 8)`LLL`, 'LAA', 'PAA', 'ALA', 'APA', 'AAL', 'AAP', 'AAA'
# n = 4, (81 - 38)
class Solution(object):
    def checkRecord(self, n):
        """
        :type n: int
        :rtype: int
        """
        power = 3 ** n

        sub = 0
        for i in range(2, n + 1):
            sub += 1
        return power - sub

if __name__ == '__main__':
    print Solution().checkRecord(4) # 43
    print Solution().checkRecord(3) # 19
    print Solution().checkRecord(2) # 8
    print Solution().checkRecord(1) # 3

