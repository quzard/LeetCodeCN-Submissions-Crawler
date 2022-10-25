/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    mid := slow.Next
    slow.Next = nil
    return merge(sortList(head), sortList(mid))
}

func merge(head1, head2 *ListNode) *ListNode {
    newHead := &ListNode{}
    cur := newHead
    for head1 != nil && head2 != nil {
        num1, num2 := head1.Val, head2.Val
        if num1 < num2 {
            head1 = head1.Next
            cur.Next = &ListNode{Val:num1}
        } else {
            head2 = head2.Next
            cur.Next = &ListNode{Val:num2}
        }
        cur = cur.Next
    }
    if head1 != nil {
        cur.Next = head1
    }
    if head2 != nil {
        cur.Next = head2
    }
    return newHead.Next
}