/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    h := head
    var tail *ListNode
    var dfs func(l *ListNode, i int)
    dfs = func(l *ListNode, i int){
        if i == 0{
            h.Next = l
            h = l
            tail = l.Next
        }else{
            dfs(l.Next, i-1)
            h.Next = l
            h = l
        }
    }
    if left == 1{
        head = &ListNode{Next:head}
        h = head
        dfs(h.Next, right - left)
        h.Next = tail
        return head.Next
    }else{
        for i:=1; i < left - 1; i++{
            h = h.Next
        }
        dfs(h.Next, right - left)
        h.Next = tail
        return head
    }
    
}