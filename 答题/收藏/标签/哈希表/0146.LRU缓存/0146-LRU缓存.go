type LRUCache struct {
    capacity int
    size     int
    head     *Node
    tail     *Node
    m     map[int]*Node
}


func Constructor(capacity int) LRUCache {
    head := &Node{}
    tail := &Node{}
    head.Next = tail
    tail.Pre = head
    return LRUCache{
        capacity: capacity,
        size: 0,
        head: head,
        tail: tail,
        m: map[int]*Node{},
    }
}


func (this *LRUCache) Get(key int) int {
    if _, ok := this.m[key]; ok {
        value := this.m[key].Val
        this.MoveToHead(this.m[key])
        return value
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.m[key]; ok {
        this.m[key].Val = value
        this.MoveToHead(this.m[key])
    } else if this.size < this.capacity {
        this.size++
        node := &Node{Val : value, Key: key}
        this.m[key] = node
        this.AddToHead(node)
    } else {
        node := &Node{Val : value, Key: key}
        this.m[key] = node
        this.AddToHead(node)
        this.RemoveTail()
    }
}

func (this *LRUCache) MoveToHead(node *Node) {
    this.DeleteNode(node)
    this.AddToHead(node)
}

func (this *LRUCache) AddToHead(node *Node) {
    this.head.Next, node.Pre, node.Next, this.head.Next.Pre = node, this.head, this.head.Next, node
}

func (this *LRUCache) DeleteNode(node *Node) {
    node.Pre.Next, node.Next.Pre = node.Next, node.Pre
}

func (this *LRUCache) RemoveTail() {
    node := this.tail.Pre
    this.DeleteNode(node)
    delete(this.m, node.Key)
}

type Node struct {
    Val int
    Key int
    Next *Node
    Pre *Node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */