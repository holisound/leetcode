class Solution(object):
    def videoStitching(self, clips, T):
        """
        :type clips: List[List[int]]
        :type T: int
        :rtype: int
        """
        clips.sort(key=lambda x: (x[0], x[1]))
        n = len(clips)
        index = 0
        # 第一个clip, start必须从0开始
        if clips[index][0] != 0:
            return -1
        # clip的end越大越好
        for i in range(1, n):
            if clips[i - 1][0] == clips[i][0]:
                index += 1
            else:
                break
        r = 0
        while index < n:
            s, e = clips[index]
            r += 1
            if e >= T:  # 已找到所有需要的clips
                return r
            end = e
            find = False
            for j in range(index + 1, n):
                c = clips[j]
                if c[0] <= e and c[1] > end:
                    # 查找最优元素，下一个clip的start小于等于当前clip的end, 并且end越大越好
                    index = j
                    end = c[1]
                    find = True
            if not find:
                return -1
        return r
