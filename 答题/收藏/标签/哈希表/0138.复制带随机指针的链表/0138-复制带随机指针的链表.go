/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList1(head *Node) *Node {
    newHead := &Node{}
    curNew := newHead
    cur := head
    hashTable := map[*Node]*Node{}
    for cur != nil{
        curNew.Next = &Node{Val: cur.Val}
        hashTable[cur] = curNew.Next
        curNew = curNew.Next
        cur = cur.Next
    }
    curNew = newHead
    cur = head
    for cur != nil{
        if cur.Random != nil{
            curNew.Next.Random = hashTable[cur.Random]
        }
        curNew = curNew.Next
        cur = cur.Next
    }
    return newHead.Next
}

var cachedNode map[*Node]*Node

func deepCopy(node *Node) *Node {
    if node == nil {
        return nil
    }
    if n, has := cachedNode[node]; has {
        return n
    }
    newNode := &Node{Val: node.Val}
    cachedNode[node] = newNode
    newNode.Next = deepCopy(node.Next)
    newNode.Random = deepCopy(node.Random)
    return newNode
}

func copyRandomList2(head *Node) *Node {
    cachedNode = map[*Node]*Node{}
    return deepCopy(head)
}



func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    for node := head; node != nil; node = node.Next.Next {
        node.Next = &Node{Val: node.Val, Next: node.Next}
    }
    for node := head; node != nil; node = node.Next.Next {
        if node.Random != nil {
            node.Next.Random = node.Random.Next
        }
    }
    headNew := head.Next
    for node := head; node != nil; node = node.Next {
        nodeNew := node.Next
        node.Next = node.Next.Next
        if nodeNew.Next != nil {
            nodeNew.Next = nodeNew.Next.Next
        }
    }
    return headNew
}
