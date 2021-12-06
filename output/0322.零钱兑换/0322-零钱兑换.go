// 背包问题
func coinChange1(coins []int, amount int) int {
    dp := make([]int, amount+1)
	dp[0] = 0
	n := len(coins)

    for i:=1; i<=amount; i++{
        res := math.MaxInt64
        for j:=0; j<n; j++{
            if i == coins[j]{
                dp[i] = 1
                break
            }else if i > coins[j]{
                if dp[i - coins[j]] > 0{
                    res = min(res, dp[i - coins[j]])
                }
            }
            if res != math.MaxInt64{
                dp[i] = res + 1
            }else{
                dp[i] = -1
            }
        }
    }
    return dp[amount]
}


func min(a, b int) int {
	if a < b {
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
