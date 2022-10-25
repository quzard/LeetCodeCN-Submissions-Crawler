func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }
    if len(coins) == 0 {
        return -1
    }

    sort.Ints(coins)
    dp := make([]int, amount + 1)
    for i := 1; i <= amount; i++ {
        dp[i] = amount + 1
        for _, coin := range coins {
            if coin <= i {
                dp[i] = min(dp[i], dp[i-coin] + 1)
            }
        }
    }
    if dp[amount] == amount + 1 {
        return -1
    }
    return dp[amount]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}