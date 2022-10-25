/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    size := 1
    tail := head
    for tail.Next != nil {
        tail = tail.Next
        size++
    }
    k %= size
    tail.Next = head
    pre := head
    i := 1
    for i + k < size {
        i++
        pre = pre.Next
    }
    res := pre.Next
    pre.Next = nil
    return res
}