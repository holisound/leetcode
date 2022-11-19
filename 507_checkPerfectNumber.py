class Solution(object):

    def checkPerfectNumber(self, num):
        """
        :type num: int
        :rtype: bool
        """
        if num == 0:
            return False
        total = 0
        s = []
        i = 1
        n = num
        while num > 0:
            if i == 1:
                total += 1
                num -= 1
                s.extend([1])
            else:
                if n % i == 0:
                    v = i + n/i
                    s.extend([i, n/i])
                    total += v
                    num -= v
            i += 1
        if total == n:
            print sorted(s), total, n
        return total == n


if __name__ == '__main__':
    for i in range(1000):
        Solution().checkPerfectNumber(i)
        # print '-' * 10