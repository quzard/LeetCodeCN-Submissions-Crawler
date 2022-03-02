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
	nextHead := slow.Next
	slow.Next = nil
	return merge(sortList(head), sortList(nextHead))
}

func merge(a, b *ListNode) *ListNode {
	result := &ListNode{}
	cur := result
	for a != nil && b != nil {
		if a.Val < b.Val {
			cur.Next = a
			a = a.Next
			cur = cur.Next
		} else {
			cur.Next = b
			b = b.Next
			cur = cur.Next
		}
	}
	if a == nil {
		cur.Next = b
	}
	if b == nil {
		cur.Next = a
	}
	return result.Next
}