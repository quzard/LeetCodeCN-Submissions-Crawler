
type Head [][]int

func (s Head) Len() int {
	return len(s)
}
func (s Head) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Head) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}


func merge1(intervals [][]int) [][]int {
    sort.Sort(Head(intervals))
    res := [][]int{}
    res = append(res, intervals[0])
    for i:=1; i<len(intervals); i++{
        if intervals[i][1] >= res[len(res) - 1][1]{
            if intervals[i][0] >= res[len(res) - 1][0] && intervals[i][0] <= res[len(res) - 1][1]{
                res[len(res) - 1][1] = intervals[i][1]
                continue
            }else if intervals[i][0] <= res[len(res) - 1][0]{
                res[len(res) - 1][1] = intervals[i][1]
                res[len(res) - 1][0] = intervals[i][0]
                continue
            }
        }else if intervals[i][0] <= res[len(res) - 1][0]{
            if intervals[i][1] >= res[len(res) - 1][0] && intervals[i][1] <= res[len(res) - 1][1]{
                res[len(res) - 1][0] = intervals[i][0]
                continue
            }
        }else if intervals[i][0] >= res[len(res) - 1][0] && intervals[i][0] <= res[len(res) - 1][1] && intervals[i][1] >= res[len(res) - 1][0] && intervals[i][1] <= res[len(res) - 1][1]{
            continue
        }
        res = append(res, intervals[i])
    }
    return res
}









func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return intervals
    }

    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] <= intervals[j][0] {
            return true
        }
        return false
    })

    begin := 0
    end := intervals[0][1]
    res := [][]int{}
    for i:=1; i<len(intervals); i++ {
        if end >= intervals[i][0] {
            end = max(intervals[i][1], end)
            continue
        }
        res = append(res, []int{intervals[begin][0], end})
        begin = i
        end = intervals[i][1]
    }

    res = append(res, []int{intervals[begin][0], end})

    return res
}

func max(a , b int ) int {
    if a > b {
        return a
    }

    return b
}