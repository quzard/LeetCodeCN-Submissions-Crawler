/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int {
    res := [][]int{}
    l :=[]int{}
    var dfs func(cur *TreeNode, sum int)
    dfs = func(cur *TreeNode, sum int) {
        if cur == nil {
            return
        }
        if cur.Left == nil && cur.Right == nil {
            sum += cur.Val
            if sum == target {
                l = append(l, cur.Val)
                res = append(res, append([]int{}, l...))
                l = l[:len(l)-1]
            }
        }
        sum += cur.Val
        if cur.Left != nil {
            l = append(l, cur.Val)
            dfs(cur.Left, sum)
            l = l[:len(l)-1]
        }
        if cur.Right != nil {
            l = append(l, cur.Val)
            dfs(cur.Right, sum)
            l = l[:len(l)-1]
        }
    }
    dfs(root, 0)
    return res
}