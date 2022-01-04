/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil && l2 == nil {
        return nil
    }

    nodes1 := []int{}
    nodes2 := []int{}
    var dfs func(node1, node2 *ListNode)
    dfs = func(node1, node2 *ListNode){
        if node1 != nil && node2 != nil{
            nodes1 = append(nodes1, node1.Val)
            nodes2 = append(nodes2, node2.Val)
            dfs(node1.Next, node2.Next)
        }else if node1 != nil{
            nodes1 = append(nodes1, node1.Val)
            dfs(node1.Next, nil)
        }else if node2 != nil{
            nodes2 = append(nodes2, node2.Val)
            dfs(nil, node2.Next)
        }
    }
    dfs(l1, l2)

    head := &ListNode{}
    var next *ListNode

    var flag int
    for len(nodes1) != 0 || len(nodes2) != 0 || flag > 0 {
        num := flag
        if len(nodes1) > 0{
            num += nodes1[len(nodes1) - 1]
            nodes1 = nodes1[:len(nodes1) - 1]
        }

        if len(nodes2) > 0{
            num += nodes2[len(nodes2) - 1]
            nodes2 = nodes2[:len(nodes2) - 1]
        }
        next = head.Next
        head.Next = &ListNode{Val: num % 10, Next:next}
        flag = num / 10
    }
    return head.Next
}