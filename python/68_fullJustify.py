class Solution(object):
    def fullJustify(self, words, maxWidth):
        """
        :type words: List[str]
        :type maxWidth: int
        :rtype: List[str]
        w1+s1+w2+
        """
        j = 0
        res = []
        curWidth = 0
        count = -1  # 初始化空格数为-1, 1个单词0个空格,2个单词1个空格, n个单词n-1空格
        for i, w in enumerate(words):
            if curWidth + count + 1 + len(w) > maxWidth:
                res.append(to_line(words[j:i], maxWidth))
                curWidth = len(w)
                j = i
                count = 0
            else:
                curWidth += len(w)
                count += 1
        last = ' '.join(words[j:])
        if last:
            last += ' ' * (maxWidth - len(last))
            res.append(last)
        return res


def to_line(words, maxWidth):
    if len(words) == 0:
        raise Exception("Empty words!")
    width = sum(len(w) for w in words)
    k = len(words) - 1
    if width + k > maxWidth:
        raise Exception("Exceed maxWidth!")
    rem = maxWidth - width
    if k > 0:
        n, m = divmod(rem, k)
    else:
        n, m = 0, rem
    res = ''
    for w in words:
        res += w
        if k > 0:
            res += ' ' * n
            k -= 1
        if m > 0:
            res += ' '
            m -= 1
    while m > 0:
        res += ' '
        m -= 1
    return res
