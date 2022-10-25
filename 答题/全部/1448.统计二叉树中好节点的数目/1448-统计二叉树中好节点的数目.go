/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    if root == nil{
        return 0
    }

    res := 1
    var dfs func(cur *TreeNode, m int)
    dfs = func(cur *TreeNode, m int){
        if cur == nil{
            return
        }
        if cur.Val >= m{
            res++
        }
        dfs(cur.Left, max(m, cur.Val))
        dfs(cur.Right, max(m, cur.Val))
    }
    dfs(root.Left, root.Val)
    dfs(root.Right, root.Val)
    return res
}
func max(a, b int) int{
    if a > b{
        return a
    }
    return b
}
