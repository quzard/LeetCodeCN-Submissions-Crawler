/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }
    l := []*TreeNode{root}
    for len(l) > 0 {
        newL := []*TreeNode{}
        vals := []int{}
        for _, node := range l {
            vals = append(vals, node.Val)
            if node.Left != nil {
                newL = append(newL, node.Left)
            }
            if node.Right != nil {
                newL = append(newL, node.Right)
            }
        }
        res = append(res, vals)
        l = newL
    }
    return res
}