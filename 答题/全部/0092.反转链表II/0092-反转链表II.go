/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head == nil {
        return nil
    }
    var tail *ListNode
    newHead := &ListNode{}
    cur := newHead
    cnt := 1
    var dfs func (node *ListNode) 
    dfs = func (node *ListNode){
        if cnt == right {
            cur.Next = node
            cur = cur.Next
            tail = node.Next
            return
        }
        cnt++
        dfs(node.Next)
        cur.Next = node
        cur = cur.Next
    }
    
    for cnt != left {
        cnt++
        cur.Next = head
        head = head.Next
        cur = cur.Next
    }
    dfs(head)
    cur.Next = tail
    return newHead.Next
}