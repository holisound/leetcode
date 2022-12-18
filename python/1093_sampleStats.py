class Solution(object):
    def sampleStats(self, count):
        """
        :type count: List[int]
        :rtype: List[float]
        """
        dim, mod = count[0], 0
        mi, mx, total = float("inf"), float('-inf'), 0
        n = 0
        samples = []
        c = []
        for i in range(256):
            if count[i] > 0:
                samples.append(float(i))
                if len(c) == 0:
                    c.append(count[i])
                else:
                    c.append(c[-1] + count[i])

                n += count[i]
                if i < mi:
                    mi = i
                if i > mx:
                    mx = i
            total += count[i] * i
            if count[i] > dim:
                dim = count[i]
                mod = i
        avg = float(total) / n

        last = c[-1]
        if last % 2 == 1:
            index = bisect.bisect_left(c, last / 2)
            mid = samples[index]
        else:
            index1, index2 = bisect.bisect_left(
                c, last / 2), bisect.bisect_left(c, last / 2 + 1)
            mid = (samples[index1] + samples[index2]) / 2
        return [mi, mx, avg, mid, mod]
