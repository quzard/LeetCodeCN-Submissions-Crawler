func maxProfit(prices []int) int {
    res := 0
    for i := 0; i < len(prices)-1; i++ {
        if prices[i] < prices[i+1] {
            res += prices[i+1] - prices[i]
        }
    }
    return res
}

func max(num1, num2 int) int {
    if num1 > num2 {
        return num1
    } else {
        return num2
    }
}
func maxProfit1(prices []int) int {
    l := len(prices)
    dp := make([][]int, l)
    for i, _ := range dp {
        dp[i] = make([]int, 2)
    }
    dp[0][0], dp[0][1] = 0, -prices[0]
    mProfile := 0
    for i := 1; i < len(prices); i++ {
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
        dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
        mProfile = max(dp[i][0], mProfile)
    }
    return mProfile
}