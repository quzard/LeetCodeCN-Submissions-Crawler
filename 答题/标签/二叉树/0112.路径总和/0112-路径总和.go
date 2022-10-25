/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    var dfs func(cur *TreeNode, sum int) bool
    dfs = func(cur *TreeNode, sum int) bool {
        sum += cur.Val
        if cur.Left == nil && cur.Right == nil {
            return sum == targetSum
        }
        if cur.Left != nil {
            if dfs(cur.Left, sum) {
                return true
            }
        }
        if cur.Right != nil {
            if dfs(cur.Right, sum) {
                return true
            }
        }
        return false
    }
    return dfs(root, 0)
}