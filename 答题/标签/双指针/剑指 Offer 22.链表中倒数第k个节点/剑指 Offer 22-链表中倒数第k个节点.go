/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
    var res *ListNode
    var dfs func(cur *ListNode)
    dfs = func(cur *ListNode){
        if cur.Next == nil{
            k--
            if k == 0{
                res = cur
                k--
            }
            return
        }
        dfs(cur.Next)
        if k > 0{
            k--
        }
        if k == 0{
            res = cur
            k--
        }
    }
    dfs(head)
    return res
}