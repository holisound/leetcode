class Solution(object):
    def findNthDigit(self, n):
        """
        :type n: int
        :rtype: int
        """
        count, base = 0, 1
        while (count + base * 9 * pow(10, base -1)) < n:
            count += base * 9 * pow(10, base -1)
            base += 1
        left = n - count
        num, mod = divmod(left, base)
        num = num if mod else num - 1
        return int(str(pow(10,base-1) + num)[mod-1])

if __name__ == '__main__':
    for s in [3, 11, 10, 999]:
        print '-------%s-------' % s
        print Solution().findNthDigit(s)