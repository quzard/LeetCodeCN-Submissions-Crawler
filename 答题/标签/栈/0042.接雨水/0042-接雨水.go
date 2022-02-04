func trap1(height []int) int {
    if len(height) <= 2{
        return 0
    }
    waters := 0
    leftHeight := height[0]
    leftIndex := 0
    rightHeight := 0
    for i := 1; i < len(height); i++{
        if height[i] >= leftHeight{
            minHeight := min(leftHeight, height[i])
            water := 0
            for j := leftIndex + 1; j < i; j++{
                if minHeight - height[j] > 0{
                    water += minHeight - height[j]
                }
            }
            waters += water
            leftHeight = height[i]
            leftIndex = i
            rightHeight = 0
        } 
        if height[i] > rightHeight{
            rightHeight = height[i]
        }
        if i == len(height) - 1{
            height[leftIndex] = rightHeight
            waters += trap(height[leftIndex:])
        }
    }
    return waters
}

func min(a, b int) int{
    if a < b{
        return a
    }
    return b
}

// 单调栈
func trap(height []int) (ans int) {
    stack := []int{}
    for i, h := range height {
        for len(stack) > 0 && h > height[stack[len(stack)-1]] {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {
                break
            }
            left := stack[len(stack)-1]
            curWidth := i - left - 1
            curHeight := min(height[left], h) - height[top]
            ans += curWidth * curHeight
        }
        stack = append(stack, i)
    }
    return
}