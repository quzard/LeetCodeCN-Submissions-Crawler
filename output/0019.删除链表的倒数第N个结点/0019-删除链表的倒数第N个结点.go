/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
    now := head
    hashMap := map[int]*ListNode{}
    index := 0
    for{
        if now != nil{
            index++
            hashMap[index] = now
            now = now.Next
        }else{
            break
        }
    }
    if n < index{
        hashMap[index - n].Next = hashMap[index - n + 1].Next
    }else{
        return head.Next
    }
    return head
}




func removeNthFromEnd(head *ListNode, n int) *ListNode {
   flag := &ListNode {
       Val: -1 ,
       Next: head,
   }
   left := flag
   right := flag

    for num := 0 ;num < n;num++ {
        right = right.Next
    }
    for right.Next != nil {
        left = left.Next
        right = right.Next
    }
    left.Next = left.Next.Next
    return flag.Next
}