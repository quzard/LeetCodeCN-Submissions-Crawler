/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    dummyHead := &ListNode{Next: head}
    tail, curr := head, head.Next
    for curr != nil {
        if tail.Val <= curr.Val {
            tail = tail.Next
        } else {
            prev := dummyHead
            for prev.Next.Val <= curr.Val {
                prev = prev.Next
            }
            tail.Next = curr.Next
            curr.Next = prev.Next
            prev.Next = curr
        }
        curr = tail.Next
    }
    return dummyHead.Next
}