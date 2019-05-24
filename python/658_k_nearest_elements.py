import bisect
class Solution(object):
    def findClosestElements(self, arr, k, x):
        index = bisect.bisect_left(arr, x)
        if index == 0:
            return arr[:k]
        lo, hi = index-1, index+1 if arr[index]==x else index
        if hi >= len(arr):
            return arr[-k:]
        res = []
        n = len(arr)
        while k > 0:
            if lo >= 0 and hi < n:
                if x - arr[lo] > arr[hi] - x:
                    res.append(arr[hi])
                    hi +=1
                else:
                    res = [arr[lo]] + res
                    lo -=1
            else:
                if lo <0:
                    res.append(arr[hi])
                    hi+=1
                elif hi >= n:
                    res = [arr[lo]] + res
                    lo-=1
            k -=1
        return res
if __name__ == '__main__':

    print Solution().findClosestElements([1,2,3,3,6,6,7,7,9,9],8,8)