func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    res := nums[0]
    sum := nums[0]
    for i := 1; i < len(nums); i++ {
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