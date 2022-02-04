/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortQuick(head *ListNode) (*ListNode, *ListNode){
    if head == nil{
        return nil, nil
    }
    if head.Next == nil{
        return head, head
    }

    mid := head
    left := &ListNode{}
    l := left
    right := &ListNode{}
    r := right
    for head.Next != nil{
        if head.Next.Val >= mid.Val{
            r.Next = head.Next
            r = r.Next
        }else{
            l.Next = head.Next
            l = l.Next
        }
        head = head.Next
    }
    l.Next = nil
    r.Next = nil
    var tail *ListNode
    if left == l{
        head = mid
        tail = mid
    }else{
        head, tail = sortQuick(left.Next)
        tail.Next = mid
        tail = mid
    }
    mid.Next = nil
    if right != r{
        mid.Next, tail = sortQuick(right.Next)
    }
    return head, tail
}

func sortList1(head *ListNode) *ListNode {
    head, _ = sortQuick(head)
    return head
}



/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	nextHead := slow.Next
	slow.Next = nil
	return merge(sortList(head), sortList(nextHead))
}

func merge(a, b *ListNode) *ListNode {
	result := &ListNode{}
	cur := result
	for a != nil && b != nil {
		if a.Val < b.Val {
			cur.Next = a
			a = a.Next
			cur = cur.Next
		} else {
			cur.Next = b
			b = b.Next
			cur = cur.Next
		}
	}
	if a == nil {
		cur.Next = b
	}
	if b == nil {
		cur.Next = a
	}
	return result.Next
}
