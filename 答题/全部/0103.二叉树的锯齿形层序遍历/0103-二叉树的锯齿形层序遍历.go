/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }
    nodes := []*TreeNode{root}
    f := false
    for len(nodes) > 0 {
        newNodes := []*TreeNode{}
        l := []int{}
        for _, node := range nodes {
            if f {
                l = append([]int{node.Val}, l...)
            } else {
                l = append(l, node.Val)
            }
            if node.Left != nil {
                newNodes = append(newNodes, node.Left)
            }
            if node.Right != nil {
                newNodes = append(newNodes, node.Right)
            }
        }
        f = !f
        res = append(res, l)
        nodes = newNodes
    }
    return res
}