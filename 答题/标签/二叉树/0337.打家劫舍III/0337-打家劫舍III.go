/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
    if root == nil {
        return 0
    }
    var dfs func(cur *TreeNode) (stealed, unStealed int)
    dfs = func(cur *TreeNode) (stealed, unStealed int) {
        stealed = cur.Val
        if cur.Left == nil && cur.Right == nil {
            return
        }
        if cur.Left != nil {
            l, r := dfs(cur.Left)
            stealed += r
            unStealed = max(l, r)
        }
        if cur.Right != nil {
            l, r := dfs(cur.Right)
            stealed += r
            unStealed += max(l, r)
        }
        return
    }
    l, r := dfs(root)
    return max(l, r)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}