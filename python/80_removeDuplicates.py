# -*- coding: utf-8 -*-
# @Author: wangwh8
# @Date:   2018-09-17 11:39:18
# @Last Modified by:   wangwh8
# @Last Modified time: 2018-09-17 13:33:46
class Solution:
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        rm=0
        dup=0
        L = len(nums)
        lo = 0
        hi = L - 1
        while lo < hi:
            if nums[lo] == nums[lo+1]:
                dup += 1
                if dup >= 2:
                    rm += 1
            else:
                if dup >= 2:
                    i = lo - dup + 1
                    j = lo
                    nums[lo - dup + 1:] = nums[j:]
                    hi -= dup - 1
                    lo -= dup - 1
                dup = 0
            lo += 1
        return L - rm
if __name__ == '__main__':
    samples = [
        [1,1,1,2,2,3],
        [0,0,1,1,1,1,2,3,3],
        [1,2,3],
        [1,2,2,2],
        [1,1,1,1],
        [1,1,1,2,2,2,3,3],
        [0,0,0,0,0,1,2,2,3,3,4,4],
        [0,0,1,1,1,1,2,2,2,4],
    ]
    for s in samples:
        pos=Solution().removeDuplicates(s)
        print s[:pos]