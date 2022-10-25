/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
    if root == nil{
        return []int{}
    }
    res := []int{}
    l := []*TreeNode{root}
    for len(l) != 0{
        newL := []*TreeNode{}
        for i := 0; i < len(l); i++{
            res = append(res, l[i].Val)
            if l[i].Left != nil{
                newL = append(newL, l[i].Left)
            }
            if l[i].Right != nil{
                newL = append(newL, l[i].Right)
            }
        }
        l = newL
    }
    return res

}