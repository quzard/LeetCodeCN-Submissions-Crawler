func maxUncrossedLines1(nums1 []int, nums2 []int) int {
    res := 0
    dp := make([][]int, len(nums1)+1)
    dp[0] = make([]int, len(nums2)+1)
    for i := 1; i <= len(nums1); i++ {
        dp[i] = make([]int, len(nums2)+1)
        for j := 1; j <= len(nums2); j++ {
            if nums1[i-1] == nums2[j-1] {
                for k := 0; k < j; k++ {
                    dp[i][j] = max(dp[i][j], dp[i-1][k] + 1)
                }
            } else {
                dp[i][j] = dp[i-1][j]
            }
            res = max(res, dp[i][j])
        }
    }
    return res
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
    dp := [2][]int{}
    dp[0] = make([]int, len(nums2) + 1)
    dp[1] = make([]int, len(nums2) + 1)
    for i := 0; i < len(nums1); i++ {
        for j := 0; j < len(nums2); j++ {
            if nums1[i] == nums2[j] {
                dp[1][j+1] = max(dp[0][j+1], dp[0][j] + 1)
            } else {
                dp[1][j+1] = max(dp[1][j], dp[0][j+1])
            }
        }
        dp[0], dp[1] = dp[1], dp[0]
    }
    return dp[0][len(nums2)]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
