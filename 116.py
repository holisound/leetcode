# Definition for binary tree with next pointer.
class TreeLinkNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
        self.next = None

class Solution:
    # @param root, a tree link node
    # @return nothing
    def connect(self, root):
        if not root: 
            return        
        stack = [root]
        
        while stack:
            size = len(stack)
            for i in range(size):
                node = stack.pop(0)
                if i < size-1:
                    node.next = stack[0]
                else:
                    node.next = None
            
                if node.left: stack.append(node.left)
                if node.right: stack.append(node.right)
if __name__ == '__main__':
    root = TreeLinkNode(1)
    root.left = TreeLinkNode(2)
    root.right = TreeLinkNode(3)
    Solution().connect(root)