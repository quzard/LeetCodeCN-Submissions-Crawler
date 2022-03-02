func maxSubArray(nums []int) int {
    sum := math.MinInt32
    res := math.MinInt32
    for i := 0; i < len(nums); i++ {
        if sum + nums[i] < nums[i] {
            sum = nums[i]
        } else {
            sum += nums[i]
        }
        res = max(res, sum)
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}