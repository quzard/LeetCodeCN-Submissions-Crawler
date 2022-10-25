// 背包问题
func profitableSchemes1(n int, minProfit int, group []int, profit []int) int {
    res := 0
    dp := make([][]int, 100 + 1)

    for i := 0; i < len(group); i++{
        dp[group[i]] = append(dp[group[i]], i)
    }
    var dfs func(num, cur, profits int)
    dfs = func(num, cur, profits int){
        if num > n{
            return
        }
        if cur == len(dp){
            return
        }
        dfs(num, cur + 1, profits)
        if len(dp[cur]) > 0{
            temp := dp[cur][len(dp[cur]) - 1]
            dp[cur] = dp[cur][:len(dp[cur])- 1]
            if len(dp[cur]) > 1{
                dfs(num, cur, profits)
            }

            profits += profit[temp]
            if profits >= minProfit && num + cur <= n{
                res++
            }
            dfs(num + cur, cur, profits)
            dp[cur] = append(dp[cur], temp)
        }
    }
    if minProfit == 0{
        res = 1
    }
    dfs(0, 0, 0)
    return res % (1e9 + 7)
}
/*
    二维费用背包问题
        选择一个工作 i ，需要两个代价：
            1. 人数减掉了 group[i]，人数最多为 n
            2. 利润增加 profit[i]，最少的利润为 minProfit
     */
func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
    var MOD = int(1e9 + 7)

    // dp[j][k]：选择了 j 个员工，并且满足工作利润至少为 k 的情况下的盈利计划的总数目
    var dp = make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, minProfit + 1)
    }

    for j := 0; j <= n; j++ {
        dp[j][0] = 1
    }

    for i := 1; i <= len(group); i++ {
        var members, earn = group[i - 1], profit[i - 1]
        for j := n; j >= members; j-- {
            for k := minProfit; k >= 0; k-- {
                // 工作利润至少为 k 而不是工作利润恰好为 k
                // 所以这个是 max(0, k - earn)，而非 k - earn，下面的代码相当于：
                if k <= earn{
                    // 相当于当前的工作产生的利润完全满足第二个代价 (即已经满足工作利润至少为 k 的条件了)
                    dp[j][k] = (dp[j][k] + dp[j - members][0]) % MOD
                }else{
                    dp[j][k] = (dp[j][k] + dp[j - members][k - earn]) % MOD
                }
                // dp[j][k] = (dp[j][k] + dp[j - members][max(0, k - earn)]) % MOD
            }
        }
    }

    return dp[n][minProfit]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}