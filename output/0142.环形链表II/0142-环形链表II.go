/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    l, f := head, head
    for f != nil && f.Next != nil && f.Next.Next != nil{
        l = l.Next
        f = f.Next.Next
        if l == f{
            l = head
            for l != f{
                l = l.Next
                f = f.Next
            }
            return f
        }
    }
    return nil
}