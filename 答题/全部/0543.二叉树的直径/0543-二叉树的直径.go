/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    res := 0
    var dfs func(cur *TreeNode) int
    dfs = func(cur *TreeNode) int {
        if cur == nil {
            return 0
        }
        left := dfs(cur.Left)
        right := dfs(cur.Right)
        res = max(res, left + right)
        return 1 + max(left, right)
    }
    res = max(res, dfs(root.Left) + dfs(root.Right))
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}