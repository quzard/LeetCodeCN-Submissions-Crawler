/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    res := [][]int{}
    var dfs func(root *TreeNode, index int)
    dfs = func (root *TreeNode, index int){
        if root == nil{
            return
        }
        if index > len(res){
            res = append(res, []int{root.Val})
        }else{
            res[index - 1] = append(res[index - 1], root.Val)
        }
        dfs(root.Left, index + 1)
        dfs(root.Right, index + 1)
    }
    dfs(root, 1)
    for l, r := 0, len(res) - 1; l < r; {
        res[l], res[r] = res[r], res[l]
        l++
        r--
    }
    return res
}