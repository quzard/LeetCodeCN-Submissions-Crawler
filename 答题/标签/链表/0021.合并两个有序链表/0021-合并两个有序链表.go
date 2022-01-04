/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{}
    num1 := 0
    num2 := 0
    now := res
    for{
        if l1 == nil && l2 == nil{
            break
        }
        if l1 != nil{
            num1 = l1.Val
        }else{
            num1 = math.MaxInt32
        }

        if l2 != nil{
            num2 = l2.Val
        }else{
            num2 = math.MaxInt32
        }
        if num1 < num2{
            now.Next = &ListNode{Val: num1}
            l1 = l1.Next
        }else{
            now.Next = &ListNode{Val: num2}
            l2 = l2.Next
        }
        now = now.Next
    }
    return res.Next
}
func min(a, b int) int{
    if a < b{
        return a
    }
    return b
}