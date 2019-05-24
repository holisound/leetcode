class Solution(object):
    def reverseStr(self, s, k):
        """
        :type s: str
        :type k: int
        :rtype: str
        """
        ans = ''
        tmp = ''
        count = 0
        for c in s:
            if count < k:
                tmp = c + tmp
                count += 1
            elif count < 2 * k:
                tmp += c
                count += 1
                if count == 2 * k:
                    ans += tmp
                    count = 0
                    tmp = ''
        if tmp:
            ans += tmp
        return ans

                        
                

if __name__ == '__main__':
    print Solution().reverseStr('abcdefg', 2)