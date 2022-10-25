/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    var dfs func(cur1, cur2 *TreeNode) *TreeNode
    dfs = func(cur1, cur2 *TreeNode) *TreeNode {
        if cur1 == nil && cur2 == nil {
            return nil
        }
        if cur2 == nil {
            cur := &TreeNode{Val: cur1.Val} 
            cur.Left = dfs(cur1.Left, cur2)
            cur.Right = dfs(cur1.Right, cur2)
            return cur
        }
        if cur1 == nil {
            cur := &TreeNode{Val: cur2.Val} 
            cur.Left = dfs(cur1, cur2.Left)
            cur.Right = dfs(cur1, cur2.Right)
            return cur
        }
        cur := &TreeNode{Val: cur1.Val + cur2.Val} 
        cur.Left = dfs(cur1.Left, cur2.Left)
        cur.Right = dfs(cur1.Right, cur2.Right)
        return cur
    }
    return dfs(root1, root2)
}