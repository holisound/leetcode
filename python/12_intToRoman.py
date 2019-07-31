class Solution(object):
    def intToRoman(self, num):
        """
        :type num: int
        :rtype: str
        """
        m = {
            "I": 1,
            "V": 5,
            "X": 10,
            "L": 50,
            "C": 100,
            "D": 500,
            "M": 1000,
        }
        m = {v: k for k, v in m.items()}
        res = ''
        digit = 1
        while num:
            last = num % 10 * digit
            if last != 0:
                if last in m:
                    res = m[last] + res
                elif last + digit in m:
                    res = m[digit] + m[last + digit] + res
                else:
                    n = 0
                    while last not in m:
                        n += 1
                        last -= digit
                    res = m[last] + m[digit] * n + res
            digit *= 10
            num /= 10

        return res
