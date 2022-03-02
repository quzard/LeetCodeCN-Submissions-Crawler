/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {

    if head == nil  || head.Next == nil{
        return  head
    }

    newHead := head.Next
    head.Next = swapPairs(newHead.Next)
    newHead.Next = head

    return newHead
}

func swapPairs1(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    newHead := &ListNode{}
    cur := newHead
    var temp *ListNode
    for head != nil && head.Next != nil {
        temp = head.Next.Next
        cur.Next = head.Next
        cur = cur.Next
        cur.Next = head
        cur = cur.Next
        cur.Next = nil
        head = temp
    }
    if head != nil{
        cur.Next = head
        cur = cur.Next
        cur.Next = nil
    }
    return newHead.Next
}