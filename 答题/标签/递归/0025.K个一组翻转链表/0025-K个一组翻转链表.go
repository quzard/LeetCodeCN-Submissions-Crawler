/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup1(head *ListNode, k int) *ListNode {
    if k <= 1 {
        return head
    }
    cnt := 0
    cur := head
    for cur != nil {
        cnt++
        cur = cur.Next
    }
    if cnt < k {
        return head
    }
    newHead := &ListNode {}
    cur = newHead
    tail := head
    var dfs func(node *ListNode, k int)
    dfs = func(node *ListNode, k int) {
        if k == 1 {
            
            tail = node.Next
            cur.Next = node
            cur = cur.Next
            cur.Next = nil
            return
        }
        dfs(node.Next, k - 1)
        cur.Next = node
        cur = cur.Next
        cur.Next = nil
    }
    for cnt >= k {
        cnt -= k
        dfs(tail, k)
    }
    cur.Next = tail
    return newHead.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    hair := &ListNode{Next: head}
    pre := hair

    for head != nil {
        tail := pre
        for i := 0; i < k; i++ {
            tail = tail.Next
            if tail == nil {
                return hair.Next
            }
        }
        nex := tail.Next
        head, tail = myReverse(head, tail)
        pre.Next = head
        tail.Next = nex
        pre = tail
        head = tail.Next
    }
    return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
    prev := tail.Next
    p := head
    for prev != tail {
        nex := p.Next
        p.Next = prev
        prev = p
        p = nex
    }
    return tail, head
}