/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
    if head.Next == nil {
        return &TreeNode{Val: head.Val}
    }

    slow, fast := head, head
    pre := slow
    for fast != nil && fast.Next != nil {
        pre = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    pre.Next = nil
    cur := &TreeNode{}
    cur.Val = slow.Val
    cur.Left = sortedListToBST(head)
    cur.Right = sortedListToBST(slow.Next)
    return cur
}

func sortedListToBST1(head *ListNode) *TreeNode {
    return buildTree(head, nil)
}

func getMedian(left, right *ListNode) *ListNode {
    fast, slow := left, left
    for fast != right && fast.Next != right {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}

func buildTree(left, right *ListNode) *TreeNode {
    if left == right {
        return nil
    }

    mid := getMedian(left, right)

    root := &TreeNode{Val: mid.Val, Left: nil, Right: nil}

    root.Left = buildTree(left, mid)
    root.Right = buildTree(mid.Next, right)

    return root
}