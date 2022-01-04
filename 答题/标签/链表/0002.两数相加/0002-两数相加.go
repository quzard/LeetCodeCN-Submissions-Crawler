/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    sum := 0
    num1 := 0
    num2 := 0
    res := &ListNode{}
    now := res
    for{
        if l1 == nil && l2 == nil{
           break 
        }
        sum = sum / 10
        if l1 != nil{
            num1 = l1.Val
            l1 = l1.Next
        }else{
            num1 = 0
        }
        if l2 != nil{
            num2 = l2.Val
            l2 = l2.Next
        }else{
            num2 = 0
        }
        sum += num1 + num2
        now.Next = &ListNode{Val:  sum % 10}
        now = now.Next
    }
    if sum >= 10{
        now.Next = &ListNode{Val:  1}
    }
    return res.Next
}