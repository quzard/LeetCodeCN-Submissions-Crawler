/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric1(root *TreeNode) bool {
    if root == nil{
        return true
    }
    nodes := []*TreeNode{root}
    for len(nodes) > 0{
        newNodes1 := []*TreeNode{}
        newNodes2 := []*TreeNode{}
        for i := 0; i < len(nodes) / 2 || i == 0; i++{
            if nodes[i].Left != nil && nodes[len(nodes) - 1 - i].Right != nil && nodes[i].Left.Val == nodes[len(nodes) - 1 - i].Right.Val{
                newNodes1 = append(newNodes1, nodes[i].Left)
                newNodes2 = append(append([]*TreeNode{}, nodes[len(nodes) - 1 - i].Right), newNodes2...)
            }else if nodes[i].Left == nil && nodes[len(nodes) - 1 - i].Right == nil{
            }else{
                return false
            }
            if nodes[i].Right != nil && nodes[len(nodes) - 1 - i].Left != nil && nodes[i].Right.Val == nodes[len(nodes) - 1 - i].Left.Val{
                if nodes[i] != nodes[len(nodes) - 1 - i]{
                    newNodes1 = append(newNodes1, nodes[i].Right)
                    newNodes2 = append(append([]*TreeNode{}, nodes[len(nodes) - 1 - i].Left), newNodes2...)
                }
            }else if nodes[i].Right == nil && nodes[len(nodes) - 1 - i].Left == nil{
            }else{
                return false
            }
        }
        nodes = append(newNodes1, newNodes2...)
    }
    return true
}


func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    q := []*TreeNode{root}
    for len(q) > 0 {
        p := []*TreeNode{}
        for i := 0; i < len(q) ; i+=2 {
            j := i + 1
            if len(q) == 1 {
                j = 0
            }
            if q[i].Val != q[j].Val {
                return false
            } else {
                if (q[i].Left != nil) != (q[j].Right != nil) {
                    return false
                } else if q[i].Left != nil && q[j].Right != nil {
                    p = append(p, q[i].Left)
                    p = append(p, q[j].Right)
                }
                if (q[j].Left != nil) != (q[i].Right != nil) {
                    return false
                } else if q[j].Left != nil && q[i].Right != nil {
                    p = append(p, q[j].Left)
                    p = append(p, q[i].Right)
                }
            }
        }
        q = p
    }
    return true
}