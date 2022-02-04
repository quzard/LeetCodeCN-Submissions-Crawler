/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
    hashTable := map[*ListNode] bool{}
    for headA != nil || headB != nil{
        if headA != nil{
            if hashTable[headA] == false{
                hashTable[headA] = true
            }else{
                return headA
            }
            headA = headA.Next
        }
        if headB != nil{
            if hashTable[headB] == false{
                hashTable[headB] = true
            }else{
                return headB
            }
            headB = headB.Next
        }
    }
    return nil
}


func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }

    pa, pb := headA, headB

    for pa != pb {
        if pa == nil {
            pa = headB
        } else {
            pa = pa.Next
        }

        if pb == nil {
            pb = headA
        } else {
            pb = pb.Next
        }
    }
    return pa
}