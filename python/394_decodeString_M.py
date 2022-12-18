class Solution(object):
    def decodeString(self, s):
        """
        :type s: str
        :rtype: str
        """
        stack=[]
        
        for c in s:
            if c == ']':
                num, seed = '', ''
                while stack and stack[-1].isalpha():
                    seed = stack.pop()+seed
                stack.pop()
                while stack and stack[-1].isdigit():
                    num = stack.pop()+num
                stack.append(seed*int(num))
                continue
            # 外层的嵌套括号'['必须保留, 作为边界
            stack.append(c)
        
        return ''.join(stack)
        
        