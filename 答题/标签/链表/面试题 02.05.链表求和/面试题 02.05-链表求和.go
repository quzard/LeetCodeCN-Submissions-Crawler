/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    newHead := &ListNode{}
    cur := newHead
    sum := 0
    for l1 != nil || l2 != nil {
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        cur.Next = &ListNode{Val: sum % 10}
        sum /= 10
        cur = cur.Next
    }
    for sum > 0 {
        cur.Next = &ListNode{Val: sum % 10}
        sum /= 10
        cur = cur.Next
    }

    return newHead.Next
}