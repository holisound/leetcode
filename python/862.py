# -*- coding: utf-8 -*-
# @Author: wangwh8
# @Date:   2018-09-21 14:25:04
# @Last Modified by:   wangwh8
# @Last Modified time: 2018-09-21 14:27:26
p = [(0, 0), (1, 56), (2, 35), (3, 91), (4, 126), (5, 117)]

tmp=[]
for i in range(1, len(p)):
    a, b = p[i-1], p[i]
    if a[0] < b[0] and a[1] < b[1]:
        tmp.append(b)
print tmp
