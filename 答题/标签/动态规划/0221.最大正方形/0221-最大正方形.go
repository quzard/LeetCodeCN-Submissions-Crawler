func maximalSquare(matrix [][]byte) int {
    res := 0
    dp := make([][]int, 2)
    dp[0] = make([]int, len(matrix[0]))
    for i := 0; i < len(matrix); i++{
        dp[1] = make([]int, len(matrix[0]))
        for j := 0; j < len(matrix[0]); j++{
            if matrix[i][j] == '1'{
                dp[1][j] = 1
                if i > 0 && j > 0{
                    dp[1][j] = max(min(dp[0][j - 1], dp[0][j], dp[1][j - 1])+ 1, dp[1][j])
                }
                res = max(dp[1][j], res)
            }
        }
        dp[0] = dp[1]
    }
    return res * res
}
func max(a, b int)int{
    if a > b{
        return a
    }
    return b
}

func min(nums ...int)int{
    m := nums[0]
    for _, num := range nums {
        if m > num {
            m = num
        }
    }
    return m
}