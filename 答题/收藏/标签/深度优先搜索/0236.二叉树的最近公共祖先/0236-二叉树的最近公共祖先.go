/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var dfs func(cur *TreeNode) *TreeNode
    dfs = func(cur *TreeNode) *TreeNode{
        if cur == nil {
            return nil
        }
        if cur == p || cur == q {
            return cur
        }
        var left, right *TreeNode
        left = dfs(cur.Left)
        right = dfs(cur.Right)
        if left == nil && right == nil {
            return nil
        }
        if left != nil && right != nil {
            return cur
        }
        if left == nil && right != nil {
            return right
        }
        if left != nil && right == nil {
            return left
        }
        return nil
    }
    return dfs(root)
}