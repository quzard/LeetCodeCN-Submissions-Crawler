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
    for head != nil {
        cur.Next = head
        cur = cur.Next
        for head.Next != nil && head.Val == head.Next.Val {
            head = head.Next
        }
        head = head.Next
        cur.Next = nil
    }
    return newHead.Next
}