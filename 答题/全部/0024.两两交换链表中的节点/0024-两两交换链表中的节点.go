/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs1(head *ListNode) *ListNode {
	var dfs func(cur *ListNode, odd bool) *ListNode
	dfs = func(cur *ListNode, odd bool) *ListNode {
		if cur == nil {
			return nil
		}
		if cur.Next == nil {
			return cur
		}
		var next, temp *ListNode
		next = dfs(cur.Next, !odd)
		if odd && next != nil {
			temp = next.Next
			next.Next = cur
			cur.Next = temp
			return next
		}
		cur.Next = next
		return cur
	}
	return dfs(head, true)
}

func swapPairs(head *ListNode) *ListNode {
	dummyNode := &ListNode{0, head}
	cur := dummyNode
	for cur.Next != nil && cur.Next.Next != nil {
		cur, cur.Next, cur.Next.Next, cur.Next.Next.Next = cur.Next, cur.Next.Next, cur.Next.Next.Next, cur.Next
	}
	return dummyNode.Next
}
