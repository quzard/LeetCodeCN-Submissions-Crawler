/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList1(head *ListNode)  {
    h := head
    finish := false

    var dfs func(now *ListNode)*ListNode
    dfs = func(now *ListNode)*ListNode{
        if now.Next == nil{
            return now
        }else{
            l := dfs(now.Next)
            if l == h || h.Next == l{
                finish = true
            }
            if finish{
                return nil
            }
            temp := h.Next
            h.Next = l
            l.Next = temp
            h = temp
            now.Next = nil
            return now 
        }
    }
    dfs(head)
}


func reorderList(head *ListNode) {
	//快慢指针,慢指针每次走一步,快指针每次走两步,快指针到达尾部的时候慢指针到达链表中间
	var stack = make([]*ListNode, 0)
	var slow, fast = head, head
	for nil != fast && nil != fast.Next {
		slow = slow.Next
		fast = fast.Next.Next
	}

	for nil != slow {
		stack = append(stack, slow)
		slow = slow.Next
	}

	for slow = head; len(stack) > 0; {
		tempNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		tempNode.Next = slow.Next
		slow.Next = tempNode
		slow = tempNode.Next
	}
	slow.Next = nil
}

