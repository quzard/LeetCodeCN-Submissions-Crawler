/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {
        return nil
    }
    var dfs func(cur *ListNode) *ListNode
    dfs = func(cur *ListNode) *ListNode {
        if cur == nil {
            return nil
        }
        cur.Next = dfs(cur.Next)
        n--
        if n == 0 {
            return cur.Next
        }
        return cur
    }
    return dfs(head)
}