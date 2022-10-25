/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    newHead := &ListNode{}
    cur := newHead
    for head != nil && head.Next != nil {
        if head.Next.Val == head.Val {
            for head.Next != nil && head.Next.Val == head.Val {
                head = head.Next
            }
            head = head.Next
        } else{
            cur.Next = head
            head = head.Next
            cur = cur.Next
            cur.Next = nil
        }
    }
    if head != nil {
        cur.Next = head
        cur = cur.Next
        cur.Next = nil
    }
    return newHead.Next
}