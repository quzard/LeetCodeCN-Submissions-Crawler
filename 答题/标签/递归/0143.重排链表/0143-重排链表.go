/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return 
    }
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    stack := []*ListNode{}
    for slow.Next != nil {
        stack = append(stack, slow.Next)
        slow = slow.Next
    }
    for len(stack) > 0 {
        temp := head.Next
        head.Next = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        head.Next.Next = temp
        head = temp
    }
    head.Next = nil
}