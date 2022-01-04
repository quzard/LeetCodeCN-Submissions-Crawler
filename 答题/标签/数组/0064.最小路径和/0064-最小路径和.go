func minPathSum(grid [][]int) int {
    dp := make([][]int, len(grid))
    for i := 0; i < len(grid); i++{
        dp[i] = make([]int, len(grid[0]))
        for j := 0; j < len(grid[0]); j++{
            if j == 0 && i == 0{
                dp[i][j] = grid[i][j]
            }else if j == 0{
                dp[i][j] = grid[i][j] + dp[i - 1][j]
            }else if i == 0{
                dp[i][j] = grid[i][j] + dp[i][j - 1]
            }else {
                dp[i][j] = min(dp[i][j - 1], dp[i - 1][j]) + grid[i][j]
            }
        }
    }
    return dp[len(grid) - 1][len(grid[0]) - 1]
}

func min(a, b int) int{
    if a < b{
        return a
    }
    return b
}