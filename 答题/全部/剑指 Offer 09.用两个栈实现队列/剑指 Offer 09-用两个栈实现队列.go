type CQueue struct {
    l1 []int
    l2 []int
}


func Constructor() CQueue {
    return CQueue{
        l1 : []int{}, 
        l2 : []int{},
    }
}


func (this *CQueue) AppendTail(value int)  {
    this.l1 = append(this.l1, value)
}


func (this *CQueue) DeleteHead() int {
    if len(this.l2) == 0 {
        for len(this.l1) > 0 {
            this.l2 = append(this.l2, this.l1[len(this.l1)-1])
            this.l1 = this.l1[:len(this.l1)-1]
        }
    }
    if len(this.l2) == 0 {
        return -1
    }
    value := this.l2[len(this.l2)-1]
    this.l2 = this.l2[:len(this.l2)-1]
    return value
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */