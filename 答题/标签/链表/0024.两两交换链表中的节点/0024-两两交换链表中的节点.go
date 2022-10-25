/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    newHead := &ListNode{}
    cur := newHead
    for head != nil && head.Next != nil {
        temp := head.Next.Next
        cur.Next = head.Next
        head.Next.Next = head
        head.Next = temp
        cur = head
        head = head.Next
    }
    if head != nil {
        cur.Next = head
        cur.Next.Next = nil
    }

    return newHead.Next
}