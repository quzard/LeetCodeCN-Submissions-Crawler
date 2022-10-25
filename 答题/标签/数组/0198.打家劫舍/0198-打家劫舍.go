func rob(nums []int) int {
    dp := make([][]int, len(nums))
    res := 0

    // 0 偷 1 不偷
    for i := 0; i < len(nums); i++ {
        dp[i] = make([]int, 2)
        if i == 0 {
            dp[i][0] = nums[i]
        } else {
            dp[i][0] = dp[i-1][1] + nums[i]
            dp[i][1] = max(dp[i-1][1], dp[i-1][0])
        }    
        res = max(res, max(dp[i][0], dp[i][1]))
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}