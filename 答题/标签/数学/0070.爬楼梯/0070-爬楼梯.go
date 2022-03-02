func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    pre, cur := 1, 2
    for i := 3; i <= n; i++ {
        pre, cur = cur, cur + pre
    }
    return cur
}