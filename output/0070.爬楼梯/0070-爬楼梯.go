func climbStairs(n int) int {
    dp := []int{1,2,0}
    if n == 1{
        return 1
    }
    if n == 2{
        return 2
    }
    for i:=3; i<=n; i++{
        dp[2] = dp[0] + dp[1]
        dp[0], dp[1] = dp[1], dp[2]
    }
    return dp[2]
}