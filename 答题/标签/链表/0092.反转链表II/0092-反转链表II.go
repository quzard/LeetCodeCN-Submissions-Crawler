/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left, right int) *ListNode {
    newHead := &ListNode{}
    newHead.Next = head
    pre := newHead
    for i := 1; i <= left-1; i++ {
        pre = pre.Next
    }
    cur := pre.Next
    for i := 0; i < right-left; i++ {
        next := cur.Next
        cur.Next = next.Next
        next.Next = pre.Next
        pre.Next = next
    }
    return newHead.Next
}

func reverseBetween1(head *ListNode, left int, right int) *ListNode {
    if head == nil {
        return nil
    }
    newHead := &ListNode{}
    cur := newHead
    cnt := 1
    for cnt != left {
        cnt++
        cur.Next = head
        head = head.Next
        cur = cur.Next
    }

    var tail *ListNode
    var dfs func(node *ListNode)
    dfs = func(node *ListNode) {
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

    dfs(head)
    cur.Next = tail
    return newHead.Next
}