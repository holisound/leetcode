class Solution(object):
    def isHappy(self, n):
        """
        :type n: int
        :rtype: bool
        """
        prodset = set()
        
        while 1:
            total = 0
            while n >= 1:
                total += (n % 10) ** 2
                n /= 10
            if total in prodset:
                return False
            elif total == 1:
                return True 
            else:
                prodset.add(total)
                n = total

if __name__ == '__main__':
    print Solution().isHappy(19)