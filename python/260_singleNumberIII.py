class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        sn = sorted(nums)
        L = len(sn)
        x, y = None, None
        for i in range(0, L, 2):
            a = sn[i]
            b = sn[i + 1]
            if a != b:
                if x is None:
                    x, y = a, b
                else:
                    if y != a:
                        break
                    y = b
        return [x, y]

if __name__ == '__main__':
    samples = [
        [1,2,1,3,2,5], # [3, 5]
        [0,1,1,2], # [2, 0]
        [0,1,2,2], # [0,1]
    ]
    for s in samples:
        print Solution().singleNumber(s)