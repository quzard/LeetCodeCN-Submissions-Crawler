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
    var dfs func (cur *ListNode) *ListNode
    dfs = func (cur *ListNode) *ListNode {
        if cur.Next == nil {
            if n == 1 {
                n--
                return nil
            }
            n--
            return cur
        }
        cur.Next = dfs(cur.Next)
        if n == 1 {
            n--
            return cur.Next
        }
        n--
        return cur
    }
    return dfs(head)
}