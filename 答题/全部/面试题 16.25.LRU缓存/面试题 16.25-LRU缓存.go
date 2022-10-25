type LRUCache struct {
    capacity int
    size int
    m map[int]*Node
    head *Node
    tail *Node
}

type Node struct {
    key int
    value int
    pre *Node
    next *Node
}

func Constructor(capacity int) LRUCache {
    head := &Node{}
    tail := &Node{}
    head.next, tail.pre = tail, head
    return LRUCache{
        capacity: capacity,
        head: head,
        tail: tail,
        m: map[int]*Node{},
    }
}


func (this *LRUCache) Get(key int) int {
    if node, ok := this.m[key]; ok {
        this.MoveToHead(key)
        return node.value
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.m[key]; ok {
        this.m[key].value = value
        this.MoveToHead(key)
    }else if this.size < this.capacity {
        this.AddToHead(key, value)
    } else {
        this.AddToHead(key, value)
        node := this.tail.pre
        this.RemoveNode(node.key)
    } 
}

func (this *LRUCache) AddToHead(key, value int) {
    node := &Node{value:value, key:key}
    node.next, node.pre, this.head.next, this.head.next.pre = this.head.next, this.head, node, node
    this.m[key] = node
    this.size++
} 

func (this *LRUCache) RemoveNode(key int) {
    node := this.m[key]
    node.pre.next, node.next.pre = node.next, node.pre
    delete(this.m, key)
    this.size--
}

func (this *LRUCache) MoveToHead(key int) {
    node := this.m[key]
    this.RemoveNode(key)
    this.AddToHead(node.key, node.value)
}
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */