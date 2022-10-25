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
    if root == nil {
        return res
    }
    l := []*TreeNode{root}
    for len(l) > 0 {
        newL := []*TreeNode{}
        for _, node := range l {
            if node.Left != nil {
                newL = append(newL, node.Left)
            }
            if node.Right != nil {
                newL = append(newL, node.Right)
            }
        }
        res = append(res, l[len(l)-1].Val)
        l = newL
    }
    return res
}