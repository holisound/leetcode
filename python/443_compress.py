class Solution(object):
    def compress(self, chars):
        """
        :type chars: List[str]
        :rtype: int
        """
        # 流下了菜鸡的眼泪
        j, k = 0, 0
        N = len(chars)
        for i in range(1, N + 1):
            if i == N or chars[j] != chars[i]:
                chars[k] = chars[j]
                if i - j > 1:
                    for x in str(i - j):
                        k += 1
                        chars[k] = x
                j, k = i, k + 1
        return k
