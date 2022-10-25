/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root == nil{
        return nil
    }
    nodes := []*Node{root.Left, root.Right}
    for len(nodes) != 0 && nodes[0] != nil{
        newNodes := []*Node{}
        for i := 0; i < len(nodes); i++{
            if i < len(nodes) - 1{
                nodes[i].Next = nodes[i + 1]
            }
            newNodes = append(newNodes, nodes[i].Left)
            newNodes = append(newNodes, nodes[i].Right)
        }
        nodes = newNodes
    }
    return root
}