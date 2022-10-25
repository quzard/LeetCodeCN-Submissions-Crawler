/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
    var res *ListNode
    var dfs func (cur *ListNode)
    dfs = func(cur *ListNode){
        if cur == nil {
            return
        }
        dfs(cur.Next)
        k--
        if k == 0 {
            res = cur
        }
    }
    dfs(head)
    return res
}