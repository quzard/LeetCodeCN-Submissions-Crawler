/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// o(n), o(n)
func connect1(root *Node) *Node {
    nodes := []*Node{root}
    for len(nodes) > 0 && nodes[0] != nil{
        newNodes := []*Node{}
        for i := 0; i < len(nodes) - 1; i++{
            nodes[i].Next = nodes[i + 1]
            if nodes[i].Left != nil{
                newNodes = append(newNodes, nodes[i].Left)
            }
            if nodes[i].Right != nil{
                newNodes = append(newNodes, nodes[i].Right)
            }
        }
        if nodes[len(nodes) - 1].Left != nil{
            newNodes = append(newNodes, nodes[len(nodes) - 1].Left)
        }
        if nodes[len(nodes) - 1].Right != nil{
            newNodes = append(newNodes, nodes[len(nodes) - 1].Right)
        }
        nodes = newNodes
    }
    return root
}

// o(n) o(1)
func connect(root *Node) *Node {
    start := root
    for start != nil {
        var nextStart, last *Node
        handle := func(cur *Node) {
            if cur == nil {
                return
            }
            if nextStart == nil {
                nextStart = cur
            }
            if last != nil {
                last.Next = cur
            }
            last = cur
        }
        for p := start; p != nil; p = p.Next {
            handle(p.Left)
            handle(p.Right)
        }
        start = nextStart
    }
    return root
}