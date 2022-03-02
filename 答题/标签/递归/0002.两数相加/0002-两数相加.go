/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    add := 0
    cur := head
    for l1 != nil || l2 != nil {
        num1, num2 := 0, 0
        if l1 != nil {
            num1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            num2 = l2.Val
            l2 = l2.Next
        }
        add += num1 + num2
        cur.Next = &ListNode{Val: add % 10}
        cur = cur.Next
        add /= 10
    }
    for add > 0 {
        cur.Next = &ListNode{Val: add % 10}
        cur = cur.Next
        add /= 10
    }
    return head.Next
}