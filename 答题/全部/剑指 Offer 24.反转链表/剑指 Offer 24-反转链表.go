/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	newHead := &ListNode{}
	for head != nil {
		temp := newHead.Next
		newHead.Next = head
		head = head.Next
		newHead.Next.Next = temp
	}
	return newHead.Next
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	res := &ListNode{}
	cur := res
	var dfs func(node *ListNode)
	dfs = func(node *ListNode) {
		if node.Next == nil {
			cur.Next = node
			cur = cur.Next
			return
		}
		dfs(node.Next)
		cur.Next = node
		cur = cur.Next
	}
	dfs(head)
	cur.Next = nil
	return res.Next
}