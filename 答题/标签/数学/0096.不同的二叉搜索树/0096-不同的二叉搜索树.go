func numTrees(n int) int {
    dp := make([]int, n + 1)
    dp[0] = 1
    for i := 1; i <= n; i++ {
        for j := 0; j < i; j++ {
            // 左边j个数 右边 i-1-j 个数, 总共i个数
            dp[i] += dp[j] * dp[i-1-j]
        }
    }
    return dp[n]
}

func numTrees1(n int) int {
    C:= 1
    for i := 0; i < n; i++ {
        C = C * 2 * (2 * i + 1) / (i + 2);
    }
    return C
}