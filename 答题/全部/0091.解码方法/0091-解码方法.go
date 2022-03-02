func numDecodings(s string) int {
    if len(s) == 0 || s[0] == '0' {
        return 0
    }
    n := len(s)
    dp := make([]int, len(s)+1)
    dp[1] = 1
    dp[0] = 1
    for i := 1; i < len(s); i++ {
        if s[i] > '0' {
            dp[i+1] += dp[i]
        }
        n, _ := strconv.Atoi(s[i-1:i+1])
        if n <= 26 && n >= 10 {
            dp[i+1] += dp[i-1]
        }
    }
    return dp[n]
}