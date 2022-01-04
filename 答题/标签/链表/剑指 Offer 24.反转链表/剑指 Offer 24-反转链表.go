/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil{
        return nil
    }
    newHead := &ListNode{}
    var dfs func(now *ListNode)
    dfs = func(now *ListNode){
        if now.Next == nil{
            newHead.Next = now
            head = newHead.Next
            return
        }
        dfs(now.Next)
        now.Next = nil
        head.Next = now
        head = head.Next
    }
    dfs(head)
    return newHead.Next
}