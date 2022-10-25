func trap(height []int) int {
    if len(height) <= 2 {
        return 0
    }
    left, right := 0, len(height) - 1
    leftHeight, rightHeight := 0, 0
    waters := 0
    for left < right {
        leftHeight = max(leftHeight, height[left])
        rightHeight = max(rightHeight, height[right])
        if leftHeight < rightHeight {
            // height[leftHeight] < height[left] < height[rightHeight]
            waters += leftHeight - height[left] 
            left++
        } else {
            // height[leftHeight] > height[right] < height[rightHeight]
            waters += rightHeight - height[right] 
            right--
        }
    }
    return waters
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}