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
        isEqual := false
        for head.Next != nil && head.Val == head.Next.Val {
            head = head.Next
            isEqual = true
        }
        if isEqual {
            head = head.Next
            continue
        }
        cur.Next = head
        head = head.Next
        cur = cur.Next
        cur.Next = nil
    }
    return newHead.Next
}