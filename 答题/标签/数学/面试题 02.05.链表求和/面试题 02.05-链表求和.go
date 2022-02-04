/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    add := 0
    res := &ListNode{}
    cur := res
    for l1 != nil || l2 != nil {
        if l1 != nil {
            add += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            add += l2.Val
            l2 = l2.Next
        }
        cur.Next = &ListNode{Val: add % 10}
        cur = cur.Next
        add /= 10
    }
    if add > 0 {
        cur.Next = &ListNode{Val: add % 10}
        cur = cur.Next
        add /= 10
    }
    return res.Next
}