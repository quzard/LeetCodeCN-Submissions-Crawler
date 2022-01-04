/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome1(head *ListNode) bool {
    l := 0
    for now:=head; now!=nil; now=now.Next{
        l++
    }
    mid := int(l / 2)
    if l % 2 == 0{
        mid--
    }
    var dfs func(root *ListNode,index int) *ListNode
    res := true
    dfs =  func(root *ListNode,index int)  *ListNode{
        if index == mid{
            if l % 2 == 0{
                if root.Val == root.Next.Val{
                    return root.Next.Next
                }else{
                    res = false
                    return nil
                }
            }else{
                return root.Next
            }
        }else{
            node := dfs(root.Next, index+1)
            if node == nil{
                return nil
            }
            if root.Val == node.Val{
                return node.Next
            }else{
                res = false
                return nil
            }
        }
    }
    fmt.Println(l, mid)
    dfs(head, 0)
    return res
}



func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil{
        return true
    }
    fast := head
    slow := head
    var pre *ListNode
    for fast != nil && fast.Next!=nil{
        fast=fast.Next.Next

        temp := slow.Next
        slow.Next = pre
        pre = slow
        slow = temp
    }
    if fast != nil {
        slow = slow.Next
    }

    for slow != nil && pre != nil{
        if slow.Val != pre.Val{
            return false
        }
        slow = slow.Next
        pre = pre.Next
    }
    return true
}