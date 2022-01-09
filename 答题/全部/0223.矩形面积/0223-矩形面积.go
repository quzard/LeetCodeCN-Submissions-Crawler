func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
    res := (ax2 - ax1) * (ay2 - ay1) + (bx2 - bx1) * (by2 - by1)
    left, right, head, end := max(ax1, bx1), min(ax2, bx2), min(ay2, by2), max(ay1, by1)
    if right > left && head > end{
        res -= (right - left) * (head - end)
    }
    
    return res 
}

func min(a, b int) int{
    if a > b {
        return b
    }
    return a
}

func max(a, b int) int{
    if a > b {
        return a
    }
    return b
}