/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    sum := 0
    var dfs func(cur *TreeNode, num int)
    dfs = func(cur *TreeNode, num int) {
        if cur.Left == nil && cur.Right == nil {
            sum += num*10 + cur.Val
            return
        }
        if cur.Left != nil {
            dfs(cur.Left, num*10+cur.Val)
        }
        if cur.Right != nil {
            dfs(cur.Right, num*10+cur.Val)
        }
    }
    dfs(root, 0)
    return sum
}