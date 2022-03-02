/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    res := &ListNode{}
    cur := res
    var dfs func(node *ListNode)
    dfs = func (node *ListNode) {
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