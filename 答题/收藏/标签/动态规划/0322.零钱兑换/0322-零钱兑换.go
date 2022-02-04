// 背包问题
func coinChange1(coins []int, amount int) int {
    if amount == 0{
        return 0
    }
    dp := make([]int, amount + 1)
    for i := 1; i <= amount; i++{
        dp[i] = math.MaxInt32
        for _, coin := range coins{
            if coin == i{
                dp[i] = 1
                break
            }else if coin < i{
                if dp[i - coin] >= 1{
                    dp[i] = min(dp[i], dp[i - coin] + 1)
                }
            }
        }
        if dp[i] == math.MaxInt32{
            dp[i] = 0
        }
    }
    if dp[len(dp) - 1] == 0{
        return -1
    }
    return dp[len(dp) - 1]
}

func min(a, b int)int{
    if a < b{
        return a
    }
    return b
}


func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	n := len(coins)

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for j := 0; j < n; j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if amount < dp[amount] {
		return -1
	}

	return dp[amount]
}