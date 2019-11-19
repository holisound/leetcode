#coding:utf-8
class Solution(object):
    def count_by_diagonal(self, m, n, k):
        if m > n:
            x, y = n, m
        else:
            x, y = m, n

        count = 0

        for i in range(1, y+1):
            if i <= x:
                count += i
            else:
                count += x - i

        # for i in range(2, x+1):
        #     if i < x:
        #         count += i
        return count

    def findKthNumber(self, m, n, k):
        """
        :type m: int
        :type n: int
        :type k: int
        :rtype: int
        """
        row = 1
        col = 1
        mm = m
        nn = n
        newk = k
        while k > 0:
            newk = k
            k -= mm + nn -1
            if k <= 0:
                break
            mm -= 1
            nn -= 1
            col += 1
            row += 1
        row_arr = [ row * i for i in range(col, n+1)]
        col_arr = [ col * i for i in range(row + 1, m+1)]
        count = 0
        ret = None
        lr = len(row_arr)
        lc = len(col_arr)
        print 'newk', newk
        for i in range(max(len(row_arr), len(col_arr))):
            for arr, length in [(row_arr, lr), (col_arr, lc)]:
                if i >= length:
                    continue
                ret = arr[i]
                count += 1
                if count >= newk:
                    break
        return ret

    def test(self, m, n, k):
        """
        乘法表矩阵，3种情况:
        1. 正方形 m = n
            1   2   3   4   5 
            2   4   6   8   10
            3   6   9   12  15
            4   8   12  16  20
            5   10  15  20  25
        2. 长方形 m > n
        3. 长方形 m < n
        """
        # m == n
        arr = []
        i = 1
        while i <= m:
            arr.append(i)
            if i > 1:
                arr.append(i)
            if i > 2:
                row = 1
                col = i
                while col > 2:
                    row += 1
                    col -= 1
                    prod = row * col
                    # if prod > i + 1:
                    #     arr.extend([i + 1, i + 1])
                    arr.append(prod)
            i += 1
        return arr


if __name__ == '__main__':
    s= Solution()
    print s.test(4,4,5)
    print s.test(5,5,5)
    # print s.findKthNumber(3,3,5)
    # print s.findKthNumber(9895, 28405, 100787757) # 31666344
    # print s.findKthNumber(2,3,6)
    # print 'near', s.count_by_diagonal(3,3,5)
    # print 'near', s.count_by_diagonal(2,3,6)
