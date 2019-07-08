# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def longestUnivaluePath(self, root):
        """
        :type root: TreeNode
        :rtype: int
        这题我学到了:
        关键是要理解自底向上(后序）的递归遍历的过程, 
        最典型的自底向上案例是 #104 二叉树最大深度
        """
        self.res = 0

        def postorder(root):
            if root:
                left = postorder(root.left)
                right = postorder(root.right)
                if root.left and root.val == root.left.val:
                    left += 1
                else:
                    left = 0
                if root.right and root.val == root.right.val:
                    right += 1
                else:
                    right = 0
                self.res = max(self.res, left + right)
                return max(left, right)  # 这里因为不是求最大深度， 不需要+1
            return 0
        postorder(root)
        return self.res
