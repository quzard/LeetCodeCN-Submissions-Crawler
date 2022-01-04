func maximalSquare(matrix [][]byte) int {
    res := 0
    dp := make([][]int, len(matrix))
    for i := 0; i < len(matrix); i++{
        dp[i] = make([]int, len(matrix[0]))
        for j := 0; j < len(matrix[0]); j++{
            if matrix[i][j] == '1'{
                dp[i][j] = 1
                if i > 0 && j > 0{
                    dp[i][j] = max(min(min(dp[i - 1][j - 1], dp[i - 1][j]), dp[i][j - 1])+ 1, dp[i][j])
                }
                res = max(dp[i][j], res)
            }
        }
    }
    return res * res
}
func max(a, b int)int{
    if a > b{
        return a
    }
    return b
}

func min(a, b int)int{
    if a < b{
        return a
    }
    return b
}

