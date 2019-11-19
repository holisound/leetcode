# -*- coding: utf-8 -*-
'''
Given a string s, partition s such that every substring of the partition is a palindrome.

Return all possible palindrome partitioning of s.

Example:

Input: "aab"
Output:
[
  ["aa","b"],
  ["a","a","b"]
]
'''
import collections
class Solution:
    def partition(self, s):
        que = collections.deque()
        que.append(list(s))
        res = []
        start = 0
        while que:
            chars = que.popleft()
            res.append(chars)
            for i in range(start, len(chars)):
                s=chars[i]
                for j in range(i+1, len(chars)):
                    s += chars[j]
                    if s == s[::-1]:
                        que.append(chars[:i]+[s]+chars[j+1:])
            start += 1
        
        print(res)
s = Solution()
s.partition("aabb")
s.partition("efe")
