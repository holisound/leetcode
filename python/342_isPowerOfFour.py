# coding:utf-8
class Solution(object):
    def isPowerOfFour(self, num):
        """
        :type num: int
        :rtype: bool
        """
        # count = 0
        # i = num
        # while i:
        #     i >>= 1
        #     count += 1
        # print 'count:', count
        # if count % 2 == 0 and num & (num - 1) == 0:
        #     print 'here', num
        #     return True
        # return False
        count = 0
        while count < 32:
            if num & 1 == 1:
                if num >> 1 == 0 and count %2 == 0:
                    return True
                else:
                    return False
            count += 1
            num >>= 1
        return False

if __name__ == '__main__':
    for s in range(100):
        print s, Solution().isPowerOfFour(s)