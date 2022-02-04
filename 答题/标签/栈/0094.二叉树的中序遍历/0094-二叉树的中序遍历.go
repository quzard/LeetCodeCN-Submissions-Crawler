/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    var dfs func(root *TreeNode)
    dfs = func (root *TreeNode){
        if root == nil{
            return
        }
        dfs(root.Left)
        res = append(res, root.Val)
        dfs(root.Right)
    }
    dfs(root)
    return res
}