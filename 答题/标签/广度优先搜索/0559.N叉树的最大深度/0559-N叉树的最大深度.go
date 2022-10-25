/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth1(root *Node) int {
    if root == nil {
        return 0
    }
    maxChildDepth := 0
    for _, child := range root.Children {
        if childDepth := maxDepth(child); childDepth > maxChildDepth {
            maxChildDepth = childDepth
        }
    }
    return maxChildDepth + 1
}

func maxDepth(root *Node) (ans int) {
    if root == nil {
        return
    }
    queue := []*Node{root}
    for len(queue) > 0 {
        q := queue
        queue = nil
        for _, node := range q {
            queue = append(queue, node.Children...)
        }
        ans++
    }
    return
}