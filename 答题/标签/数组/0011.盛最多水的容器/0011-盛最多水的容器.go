func maxArea(height []int) int {
    if len(height) <= 1 {
        return 0
    }
    l, r := 0, len(height) - 1
    waters := 0
    for l < r {
        waters = max(waters, (r - l) * min(height[l], height[r]))
        if height[l] < height[r] {
            l++
        } else {
            r--
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

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}