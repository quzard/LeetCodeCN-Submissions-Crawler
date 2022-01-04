func largestRectangleArea1(heights []int) int {
    res := 0
    n := len(heights)
    for i:=0; i<n; i++{
        height := heights[i]
        res = max(res, height)
        for j:=i+1; j<n && (n-1-i+1)*height>res;j++{
            height = min(height, heights[j])
            res = max(res, (j-i+1)*height)
        }
    }
    return res
}

func largestRectangleArea(heights []int) int {
    heights = append([]int{0}, heights...)
    heights = append(heights, 0)
    var minStack = make([]int, len(heights))
    var max int
    minStackIndex := 0
    minStack[minStackIndex] = 0
    max = heights[0]
    var tmp int
    for index := 1; index < len(heights); index++ {
        if heights[index] >= heights[minStack[minStackIndex]] {
            minStackIndex++
            minStack[minStackIndex] = index
            continue
        }

        right := minStack[minStackIndex]
        var i int
        for i = minStackIndex; i >= 0; i-- {
            if heights[index] >= heights[minStack[i]] {
                break
            }
            if i > 0 {
                tmp = (right - minStack[i-1]) * heights[minStack[i]]
            } else {
                tmp = (right - minStack[i] + 1) * heights[minStack[i]]
            }
            if tmp > max {
                max = tmp
            }
        }

        minStackIndex = i+1
        minStack[minStackIndex] = index
    }

    return max
}


func max(a, b int) int{
    if a > b{
        return a
    }
    return b
}

func min(a, b int) int{
    if a < b{
        return a
    }
    return b
}