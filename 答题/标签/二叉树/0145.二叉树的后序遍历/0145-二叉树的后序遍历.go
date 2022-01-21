/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    res := []int{}
    var dfs func(cur *TreeNode)
    dfs = func(cur *TreeNode) {
        if cur == nil {
            return
        }
        if cur.Left == nil && cur.Right == nil {
            res = append(res, cur.Val)
            return
        }
        dfs(cur.Left)
        dfs(cur.Right)
        res = append(res, cur.Val)
    }
    dfs(root)
    return res
}