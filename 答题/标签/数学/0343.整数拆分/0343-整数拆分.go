func integerBreak1(n int) int {
    dp := make([]int, n + 1)
    dp[0], dp[1] = 1, 1
    for i := 2; i <= n; i++ {
        if i < n {
            dp[i] = i
        }
        for j := 1; j < i; j++ {
            dp[i] = max(dp[i], dp[i - j] * dp[j])
        }
    }
    return dp[n]
}

func max(a, b int) int{
    if a > b {
        return a
    }
    return b
}

func integerBreak(n int) int {
    if n < 4 {
        return n - 1
    }
    dp := make([]int, n + 1)
    dp[2] = 1
    for i := 3; i <= n; i++ {
        dp[i] = max(max(2 * (i - 2), 2 * dp[i - 2]), max(3 * (i - 3), 3 * dp[i - 3]))
    }
    return dp[n]
}
