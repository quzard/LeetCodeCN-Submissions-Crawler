/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode)int{
    if root == nil{
        return 0
    }
    left := max(dfs(root.Left), 0)
    right := max(dfs(root.Right), 0)
    res = max(res, left + right + root.Val)

    return root.Val+max(left, right)
}

var res int

func maxPathSum(root *TreeNode) int {
    res = math.MinInt32
    res = max(res, dfs(root))
    return res
}

func max(a, b int) int{
    if a < b{
        return b
    }
    return a
}
