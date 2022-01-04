func rob1(nums []int) int {
    if len(nums) == 0{
        return 0
    }
    res := nums[0]
    dp := make([][]int, len(nums))
    dp[0] = []int{nums[0], 0} // 偷nums[i], 不偷nums[i]
    for i:=1; i<len(nums);i++{
        dp[i] = []int{0, 0}
        if i > 1{
            dp[i][0] = max(max(dp[i - 1][1], dp[i - 2][0]),dp[i - 2][1]) + + nums[i]  // 偷nums[i]
            dp[i][1] = max(dp[i - 1][0], dp[i - 1][0])    // 不偷nums[i]
        }else{
            dp[i][0] = nums[i]  // 偷nums[i]
            dp[i][1] = max(dp[i - 1][0], dp[i - 1][0])    // 不偷nums[i]
        }
        res = max(max(res, dp[i][0]), dp[i][1])
    }
    fmt.Println(dp)
    return res
}

func max(a, b int) int{
    if a > b{
        return a
    }
    return b
}






func rob(nums []int) int {
    if len(nums) < 2 {
        return nums[0]
    }
    dp1 := nums[0]
    dp2 := max(nums[0], nums[1])
    for i := 2; i < len(nums); i++{
        ans := max(dp2, dp1 + nums[i])
        dp1 = dp2
        dp2 = ans
    }
    return dp2
}