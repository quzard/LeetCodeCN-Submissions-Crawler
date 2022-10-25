func splitArray(nums []int, m int) int {
    n := len(nums)
    
    f := make([][]int, n + 1)
    sub := make([]int, n + 1)
    for i := 0; i < len(f); i++ {
        f[i] = make([]int, m + 1)
        for j := 0; j < len(f[i]); j++ {
            f[i][j] = math.MaxInt32
        }
    }
    // 前缀和
    for i := 0; i < n; i++ {
        sub[i + 1] = sub[i] + nums[i]
    }
    f[0][0] = 0
    // 前 i 个数
    for i := 1; i <= n; i++ {
        // j 表示切分为 j 段
        for j := 1; j <= min(i, m); j++ {
            // 将数组的前 k 个数分割为 j - 1 段
            for k := j - 1; k < i; k++ {
                // f[i][j] 表示将数组的前 i 个数分割为 j 段所能得到的最大连续子数组和的最小值
                // f[k][j - 1] 表示将数组的前 k 个数分割为 j - 1 段所能得到的最大连续子数组和的最小值
                // sub[i] - sub[k] 表示将后（k-i) 个数作为一段的和
                f[i][j] = min(f[i][j], max(f[k][j - 1], sub[i] - sub[k]))
            }
        }
    }
    return f[n][m]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}