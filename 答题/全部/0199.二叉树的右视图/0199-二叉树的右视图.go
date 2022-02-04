/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    res := []int{}
    var dfs func(root *TreeNode, height int)
    dfs = func(root *TreeNode, height int){
        if root == nil{
            return
        }
        if len(res) < height{
            res = append(res, root.Val)
        }
        dfs(root.Right, height + 1)
        dfs(root.Left, height + 1)
    }
    dfs(root, 1)
    return res
}