func rob(nums []int) int {
    dp := [2]int{}
    // [0] 偷 [1] 不偷
    for i := 0; i < len(nums); i++ {
        if i == 0 {
            dp[0] = nums[i]
            continue
        }
        dp[0], dp[1] = dp[1] + nums[i], max(dp[0], dp[1])
    }
    return max(dp[0], dp[1])
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}