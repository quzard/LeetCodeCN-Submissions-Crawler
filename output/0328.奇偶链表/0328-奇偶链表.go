/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList1(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    slow := head
    fast := head.Next
    for fast != nil && fast.Next != nil{
        temp := slow.Next
        slow.Next = fast.Next
        slow.Next.Next, fast.Next = temp, slow.Next.Next
        slow = slow.Next
        fast = fast.Next
    }
    return head
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil{
		return nil 
	}
	if head.Next == nil{
		return head
	}
	evenhead := &ListNode{}
	odd := head
	for even:= evenhead;odd != nil && even !=nil;odd=odd.Next{
		even.Next = odd.Next
        even = even.Next
        if even == nil || even.Next == nil{
            break
        }
		odd.Next = even.Next
	}
	odd.Next = evenhead.Next
	return head
	
}