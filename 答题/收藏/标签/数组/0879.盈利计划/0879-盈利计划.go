// ��������
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
    ��ά���ñ�������
        ѡ��һ������ i ����Ҫ�������ۣ�
            1. ���������� group[i]���������Ϊ n
            2. �������� profit[i]�����ٵ�����Ϊ minProfit
     */
func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
    var MOD = int(1e9 + 7)

    // dp[j][k]��ѡ���� j ��Ա�����������㹤����������Ϊ k ������µ�ӯ���ƻ�������Ŀ
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
                // ������������Ϊ k �����ǹ�������ǡ��Ϊ k
                // ��������� max(0, k - earn)������ k - earn������Ĵ����൱�ڣ�
                if k <= earn{
                    // �൱�ڵ�ǰ�Ĺ���������������ȫ����ڶ������� (���Ѿ����㹤����������Ϊ k ��������)
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