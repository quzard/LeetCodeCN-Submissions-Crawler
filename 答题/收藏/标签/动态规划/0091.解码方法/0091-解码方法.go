func numDecodings(s string) int {
    if len(s) == 0 || s[0] == '0'{
        return 0
    }
    dp := make([]int, len(s) + 1)
    dp[len(s)] = 1
    for i := len(s) - 1; i >= 0; i--{
        if s[i] == '0'{
            continue
        }
        dp[i] += dp[i + 1]
        if i + 1 < len(s){
            if s[i:i+2] <= "26"{
                dp[i] += dp[i + 2]
            }
        }
    }
    return dp[0]  
}