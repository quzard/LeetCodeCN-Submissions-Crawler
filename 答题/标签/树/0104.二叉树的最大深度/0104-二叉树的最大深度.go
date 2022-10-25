/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    res := 0
    var dfs func(cur *TreeNode, sum int) 
    dfs = func(cur *TreeNode, sum int){
        if cur == nil {
            res = max(res, sum)
            return
        }
        dfs(cur.Left, sum + 1)
        dfs(cur.Right, sum + 1)
    }
    dfs(root, 0)
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}