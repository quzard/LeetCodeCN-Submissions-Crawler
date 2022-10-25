/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList1(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 存slow.Next 而不是slow
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

func reorderList(head *ListNode) {
	if head.Next == nil {
		return
	}
	fast, slow := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	fast = slow.Next
	slow.Next = nil
	slow = head

	pre, next := fast, fast.Next

	for next != nil {
		tmp := next.Next
		next.Next = pre
		pre = next
		next = tmp
	}
	fast.Next = nil
	fast = pre

	for fast != nil {
		tmp1, tmp2 := slow.Next, fast.Next
		slow.Next = fast
		fast.Next = tmp1
		slow, fast = tmp1, tmp2
	}
}