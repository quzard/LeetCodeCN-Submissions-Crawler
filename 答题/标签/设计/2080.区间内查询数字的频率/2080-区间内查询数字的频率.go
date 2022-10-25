type RangeFreqQuery struct {
    m map[int]sort.IntSlice
}


func Constructor(arr []int) RangeFreqQuery {
    rangeFreqQuery := RangeFreqQuery{
        m : map[int]sort.IntSlice{},
    }
    for i := 0; i < len(arr); i++ {
        rangeFreqQuery.m[arr[i]] = append(rangeFreqQuery.m[arr[i]], i)
    }
    return rangeFreqQuery
}


func (this *RangeFreqQuery) Query(left int, right int, value int) int {
    arr := this.m[value]
    return arr[arr.Search(left):].Search(right+1)
}


/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */