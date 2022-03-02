func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return [][]int{}
    }
    sort.Slice(intervals, func(a, b int) bool {
        if intervals[a][0] == intervals[b][0] {
            return intervals[a][1] < intervals[b][1]
        } 
        return intervals[a][0] < intervals[b][0]
    })
    res := [][]int{}
    l, r :=  intervals[0][0], intervals[0][1]
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] > r {
            res = append(res, []int{l, r})
            l, r = intervals[i][0], intervals[i][1]
        } else if intervals[i][0] == r {
            r = intervals[i][1]
        } else if intervals[i][0] < r && intervals[i][0] >= l && intervals[i][1] >= r {
            r = intervals[i][1]
        } 
    }
    res = append(res, []int{l, r})
    return res
}