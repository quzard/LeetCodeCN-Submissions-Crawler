type CQueue struct {
    data []int
}


func Constructor() CQueue {
    return CQueue{
        data: []int{},
    }
}


func (this *CQueue) AppendTail(value int)  {
    this.data = append(this.data, value)
}


func (this *CQueue) DeleteHead() int {
    if len(this.data) == 0{
        return -1
    }else{
        ans := this.data[0]
        this.data = this.data[1:]
        return ans
    }
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */