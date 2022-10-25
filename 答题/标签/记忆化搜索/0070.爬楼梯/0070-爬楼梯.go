func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    dp := []int{1,1,0}
    for i := 2; i <= n; i++ {
        dp[2] = dp[1] + dp[0]
        dp[1], dp[0] = dp[2], dp[1]
    }
    return dp[2]
}