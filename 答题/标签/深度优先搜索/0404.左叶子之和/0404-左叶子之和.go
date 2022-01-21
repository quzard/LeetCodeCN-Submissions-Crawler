/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    res := 0
    var dfs func(cur *TreeNode, left bool)
    dfs = func(cur *TreeNode, left bool) {
        if cur == nil {
            return
        }
        if cur.Left == nil && cur.Right == nil && left {
            res += cur.Val
            return
        }
        dfs(cur.Left, true)
        dfs(cur.Right, false)
    }
    dfs(root, false)
    return res
}