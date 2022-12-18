class Solution:
    def minOperations(self, boxes: str) -> List[int]:
        ans = []
        left = right = total = 0
        n = len(boxes)
        for i in range(n):
            if boxes[i] == '1':
                right += 1
                total += i
        for i in range(n):
            ans.append(total)
            if boxes[i] == '1':
                left += 1
                right -= 1
            total += left - right
        return ans
# 第二种解法似曾相识（强人太多了😂）
class Solution:
    def minOperations(self, boxes: str) -> List[int]:
        res = []
        n = len(boxes)
        now = total = 0
        for i in range(n):
            total += now
            res.append(total)
            if boxes[i] == '1':
                now += 1
        now = total = 0
        for i in range(n-1,-1,-1):
            total += now
            res[i] += total
            if boxes[i] == '1':
                now += 1
        return res
