func maxArea(height []int) int {
    if len(height) <= 1 {
        return 0
    }
    l, r := 0, len(height) - 1
    water := 0
    for l < r {
        water = max(water, (r - l) * min(height[l], height[r]))
        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }
    return water
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