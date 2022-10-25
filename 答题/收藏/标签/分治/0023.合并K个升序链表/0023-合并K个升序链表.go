/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    if len(lists) == 1 {
        return lists[0]
    }
    mid := len(lists) / 2
    left := mergeKLists(lists[:mid])
    right := mergeKLists(lists[mid:])
    newhead := &ListNode{}
    cur := newhead
    for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
            left = left.Next
		} else {
			cur.Next = right
            right = right.Next
		}
        cur = cur.Next
    }
    if left != nil {
        cur.Next = left
    } 
    if right != nil {
        cur.Next = right
    }
    return newhead.Next
}
