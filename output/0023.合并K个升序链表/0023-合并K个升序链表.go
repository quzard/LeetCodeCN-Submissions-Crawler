/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists1(lists []*ListNode) *ListNode {
    head := &ListNode{}
    now := head
    for{
        var node *ListNode
        var index int
        isExsited := false
        for i, l := range(lists){
            if l != nil{
                isExsited = true
                if node == nil{
                    node = l
                    index = i
                }else{
                    if node.Val > l.Val{
                        index = i
                        node = l
                    }
                }
            }
            
        }
        if isExsited == false{
            break
        }
        now.Next = &ListNode{
            Val: node.Val,
        }
        now = now.Next
        lists[index] = lists[index].Next
    }
    return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	leftNode := mergeKLists(lists[:mid])
	rightNode := mergeKLists(lists[mid:])
	dumpHead := &ListNode{}
	pre := dumpHead
	for leftNode != nil && rightNode != nil {
		if leftNode.Val < rightNode.Val {
			pre.Next, pre, leftNode = leftNode, leftNode, leftNode.Next
		} else {
			pre.Next, pre, rightNode = rightNode, rightNode, rightNode.Next
		}
	}
	if leftNode == nil {
		pre.Next = rightNode
	}
	if rightNode == nil {
		pre.Next = leftNode
	}
	return dumpHead.Next
}