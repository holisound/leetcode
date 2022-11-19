# coding:utf-8
# 示例 1：

# 输入："{a,b}c{d,e}f"
# 输出：["acdf","acef","bcdf","bcef"]
# 示例 2：

# 输入："abcd"
# 输出：["abcd"]


def run(s):
    i, n = 0, len(s)
    res = ['']
    while i < n:
        c = s[i]
        if c != '{':
            res = [x + c for x in res]
            i += 1
        else:
            temp = []
            for j in range(i + 1, n):
                if s[j] == '}':
                    i = j + 1
                    break
                if s[j] != ',':
                    for x in res:
                        temp.append(x + s[j])
            res = temp

    return res


if __name__ == '__main__':
    print run("{a,b}c{d,e}f")
    print run("abcd")
