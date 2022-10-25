func maximalSquare(matrix [][]byte) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }
    res := 0
    dp := make([][]int, len(matrix))
    for i := 0; i < len(matrix); i++ {
        dp[i] = make([]int, len(matrix[0]))
        for j := 0; j < len(matrix[0]); j++ {
            if matrix[i][j] == '0' {
                continue
            }
            if i == 0 || j == 0 {
                dp[i][j] = 1
            } else {
                dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
            }
            res = max(res, dp[i][j])
        }
    }

    return res * res
}

func min(nums ...int) int {
    m := math.MaxInt
    for _, num := range nums {
        if m > num {
            m = num
        }
    }
    return m
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}