/*
https://leetcode.cn/problems/linked-list-in-binary-tree/
1367. 二叉树中的链表
*/
func isSubPath(head *ListNode, root *TreeNode) bool {
    if root == nil || head == nil{
        return false
    }
    var res bool
    if root.Val == head.Val {
        res = dfs(head, root)
    }
    return res || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dfs(head *ListNode, root *TreeNode) bool{
    if head == nil {
        return true
    }
    if root == nil {
        return false
    }
    return head.Val == root.Val && (
         dfs(head.Next, root.Left) || dfs(head.Next, root.Right))
}
