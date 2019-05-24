class Solution(object):
    def largestNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: str
        """
        max_num = max(nums)
        if max_num == 0:
            return "0"
        max_len = 1
        while max_num >= 10:
            max_num /= 10
            max_len += 1

        def sortfn(n):
            max_factor = 0
            ilen = 1
            i = n
            while i >= 10:
                max_factor = max(max_factor, i % 10)
                i /= 10
                ilen += 1
            for _ in range(max_len - ilen):
                n = n * 10 + max(max_factor, i)
            return n

        ret = ''.join(str(n) for n in sorted(nums, key=sortfn, reverse=True))
        return ret
                
if __name__ == '__main__':
    print Solution().largestNumber([3,610,614,5,9, 61, 6])#"966161461053"
    print Solution().largestNumber([3,30,34,5,9]) #9534330
    print Solution().largestNumber([10, 2]) #9534330
    print Solution().largestNumber([121, 12]) #9534330