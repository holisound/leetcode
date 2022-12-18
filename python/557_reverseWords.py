import timeit
class Solution(object):
    def reverseWords(self, s):
        """
        :type s: str
        :rtype: str
        """
        return ' '.join(i[::-1] for i in s.split())

    def reverseWordsV2(self, s):
        L = len(s)
        word = ''
        tmp = []
        for i in range(L):
            w = s[i]
            if w == ' ':
                tmp.append(word)
                word = ''
            else:
                word = w + word
        if word:
            tmp.append(word)
        return ' '.join(tmp)

if __name__ == '__main__':
    print timeit.timeit('Solution().reverseWords("Let\'s take LeetCode contest")', "from __main__ import Solution", number=1000)
    print timeit.timeit('Solution().reverseWordsV2("Let\'s take LeetCode contest")', "from __main__ import Solution", number=1000)
