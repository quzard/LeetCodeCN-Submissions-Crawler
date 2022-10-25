/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup( head *ListNode ,  k int ) *ListNode {
    if k == 1 {
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
    
    newHead := &ListNode{}
    cur = newHead
    start := head
    var dfs func(node *ListNode, k int)
    dfs = func(node *ListNode, k int) {
        if k == 0 {
            start = node
            return
        }
        dfs(node.Next, k-1)
        cur.Next = node
        cur = cur.Next
        cur.Next = nil
    }
    for cnt >= k {
        cnt -= k
        dfs(start, k)
    }
    cur.Next = start
    return newHead.Next
}



func reverseKGroup1(head *ListNode, k int) *ListNode {
    newHead := &ListNode{Next: head}
    cur := newHead

    for head != nil {
        tail := head
        // 此时tail为第k个节点
        for i := 1; i < k; i++ {
            tail = tail.Next
            if tail == nil {
                return newHead.Next
            }
        }
        //nextHead为下一个
        nextHead := tail.Next
        //反转第1个节点~第k个节点
        head, tail = myReverse(head, tail)

        cur.Next = head
        tail.Next = nextHead

        cur = tail
        head = nextHead
    }
    return newHead.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
    prev := tail.Next
    cur := head
    for prev != tail {
        next := cur.Next

        cur.Next = prev

        prev = cur

        cur = next
    }
    return tail, head
}