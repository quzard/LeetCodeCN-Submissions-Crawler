/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
    res := [][]int{}
    var dfs func(cur *TreeNode, sum int, l []int)
    dfs = func(cur *TreeNode, sum int, l []int){
        if cur == nil{
            return
        }
        if cur.Left == nil && cur.Right == nil{
            l = append(l, cur.Val)
            if sum + cur.Val == targetSum{
                res = append(res, append([]int{}, l...))
            }
            l = l[:len(l) - 1]
            return
        }
        l = append(l, cur.Val)
        dfs(cur.Left, sum + cur.Val, l)
        dfs(cur.Right, sum + cur.Val, l)
        l = l[:len(l) - 1]
    }
    dfs(root, 0, []int{})
    return res
}