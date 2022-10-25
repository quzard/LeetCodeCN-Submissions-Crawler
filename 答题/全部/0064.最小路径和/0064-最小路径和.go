func minPathSum(grid [][]int) int {
    dp := make([][]int, 2)
    dp[0] = make([]int, len(grid[0]))
    dp[1] = make([]int, len(grid[0]))
    for i := 0; i < len(grid); i++{
        for j := 0; j < len(grid[0]); j++{
            if j == 0 && i == 0{
                dp[1][j] = grid[i][j]
            }else if j == 0{
                dp[1][j] = grid[i][j] + dp[0][j]
            }else if i == 0{
                dp[1][j] = grid[i][j] + dp[1][j - 1]
            }else {
                dp[1][j] = min(dp[1][j - 1], dp[0][j]) + grid[i][j]
            }
        }
        dp[1], dp[0] = dp[0], dp[1]
    }
    return dp[0][len(grid[0]) - 1]
}

func min(a, b int) int{
    if a < b{
        return a
    }
    return b
}