/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    var dfs func(left, right *TreeNode) bool
    dfs = func(left, right *TreeNode) bool {
        if left == nil && right == nil {
            return true
        }
        if left == nil || right == nil {
            return false
        }
        if left.Val != right.Val {
            return false
        }
        return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
    }
    return dfs(root, root)
}