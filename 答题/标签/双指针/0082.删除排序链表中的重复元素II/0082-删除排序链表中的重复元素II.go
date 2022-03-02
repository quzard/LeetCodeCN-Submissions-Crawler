/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    res := &ListNode{}
    cur := res
    for head != nil {
        if head.Next != nil && head.Next.Val == head.Val{
            val := head.Val
            for head != nil && head.Val == val {
                head = head.Next
            }
        } else {
            cur.Next = &ListNode{Val: head.Val}
            cur = cur.Next
            head = head.Next
        }
    }
    return res.Next
}