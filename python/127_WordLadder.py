from heapq import *

class Solution:
    def ladderLength(self, beginWord: str, endWord: str, wordList: List[str]) -> int:
        # 优先队列（最小堆）
        # BFS
        wordSet=set(wordList)
        que=[(1, beginWord)]
        while que:
            cnt,w = heappop(que)
            if w==endWord:
                return cnt
            for i in range(len(w)):
                for c in string.ascii_lowercase:
                    s=w[:i]+c+w[i+1:]
                    if s in wordSet:
                        wordSet.remove(s)
                        que.append((cnt+1,s))
        return 0
