/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    res := 0
    var dfs func(cur *TreeNode)int
    dfs = func(cur *TreeNode)int{
        if cur == nil{
            return 0
        }
        left := dfs(cur.Left)
        right := dfs(cur.Right)
        res = max(left + right, res)
        return max(left, right) + 1
    }
    dfs(root)
    return res
}

func max(a, b int) int{
    if a > b {
        return a
    }
    return b
}