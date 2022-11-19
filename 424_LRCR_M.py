class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        win, j, res = {}, 0, 0

        for i, c in enumerate(s):
            win[c] = win.get(c, 0) + 1
            most_freq_c = max(win, key=win.get)
            if i - j + 1 - win[most_freq_c] > k:
                win[s[j]] -= 1
                j += 1
            res = max(res, i - j + 1)
        return res
    
    def characterReplacement2(self, s, k):
        maxf = res = 0
        count = collections.Counter()
        for i in range(len(s)):
            count[s[i]] += 1
            maxf = max(maxf, count[s[i]])
            if res - maxf < k: res += 1
            else: count[s[i - res]] -= 1
        return res
