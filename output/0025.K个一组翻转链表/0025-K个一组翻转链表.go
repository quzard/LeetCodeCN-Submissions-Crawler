/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    newHead := &ListNode{Next: &ListNode{}}
    var tail *ListNode
    tail = nil
    cur := head
    head = newHead
    i := k
    for {
        if cur == nil && i == 0{
            break
        }
        if cur == nil && i > 0 {
            tail.Next = nil
            tail = nil
            cur = head.Next
            i = k - i
        }
        if i == 0{
            i = k
            head = tail
            tail = nil
        }
        temp := head.Next 
        next := cur.Next
        head.Next = cur
        cur.Next = temp
        cur = next
        i--
        if tail == nil{
            tail = head.Next
        }
    }
    tail.Next = nil
    return newHead.Next
}
