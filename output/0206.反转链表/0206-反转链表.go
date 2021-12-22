/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    res := &ListNode{}
    for{
        if head == nil{
            break
        }
        temp := res.Next
        res.Next = &ListNode{Val: head.Val}
        res.Next.Next = temp
        head = head.Next
    }

    return res.Next
}