class Solution(object):
    def compareVersion(self, version1, version2):
        """
        :type version1: str
        :type version2: str
        :rtype: int
        """
        if version1 == version2:
            return 0
        s1 = [int(i) for i in version1.split('.')]
        s2 = [int(i) for i in version2.split('.')]
        while len(s1) != len(s2):
            if len(s1) > len(s2):
                s2.append(0)
            else:
                s1.append(0)
        if len(s1) == len(s2) == 1:
            if s1[0] == s2[0]:
                return 0
            elif s1[0] > s2[0]:
                return 1
            else:
                return -1
        def redu(x,y):
            return x * 10 + y
        r1 = reduce(redu, s1)
        r2 = reduce(redu, s2)
        if r1 > r2:
            return 1
        elif r1 == r2:
            return 0
        return -1
                

    def compareVersion2(self, version1, version2):
        """
        :type version1: str
        :type version2: str
        :rtype: int
        """
        v1_split = version1.split('.')
        v2_split = version2.split('.')
        v1_len, v2_len = len(v1_split), len(v2_split)
        maxLen = max(v1_len, v2_len)
        for i in range(maxLen):
            temp1, temp2 = 0, 0
            if i < v1_len:
                temp1 = int(v1_split[i])
            if i < v2_len:
                temp2 = int(v2_split[i])
            if temp1 < temp2:
                return -1
            elif temp1 > temp2:
                return 1
        return 0
if __name__ == '__main__':
    s = Solution()
    import time
    a = time.clock()
    print s.compareVersion("1.9", "1.9.9")
    b = time.clock()
    print s.compareVersion2("1.9", "1.9.9")
    c = time.clock()
    print b-a
    print c-b