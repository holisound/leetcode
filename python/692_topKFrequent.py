from Queue import PriorityQueue
class Solution(object):
    def topKFrequent(self, words, k):
        """
        :type words: List[str]
        :type k: int
        :rtype: List[str]
        """
        #优先队列
        count = collections.Counter(words)
        que = PriorityQueue()
        for j in count:
            que.put((-count[j], j))
        ans=[]
        while k:
            ans.append(que.get()[1])
            k -= 1
            
        return ans
            
            
        