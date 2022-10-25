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

func largestRectangleArea2(heights []int) int {
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

// 单调栈
func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
        // 说明 0~i-1 都比 i 大
		if len(stack) == 0 {
			left[i] = -1
		} else {
            // 说明 stack[len(stack)-1] + 1 ~ i - 1 都比 i 大
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
        // 说明 i+1~n-1 都比 i 大
		if len(stack) == 0 {
			right[i] = n
		} else {
             // 说明 i+1~stack[len(stack)-1] 都比 i 大
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
    // 面积= (right[i]-left[i]-1)*heights[i])
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}