/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    var dfs func(cur *TreeNode)
    dfs = func(cur *TreeNode){
        if cur == nil{
            return
        }
        cur.Right, cur.Left = cur.Left, cur.Right
        dfs(cur.Right)
        dfs(cur.Left)
    }
    dfs(root)
    return root
}