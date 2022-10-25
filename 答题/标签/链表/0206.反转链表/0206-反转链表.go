/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    newHead := &ListNode{}
    cur := newHead
    var dfs func(n *ListNode)
    dfs = func(n *ListNode) {
        if n == nil {
            return
        }
        dfs(n.Next)
        cur.Next = n
        cur = cur.Next
        cur.Next = nil
    }
    dfs(head)
    return newHead.Next
}