func maxSubArray(nums []int) int {
    if len(nums) == 0{
        return 0
    }
    res := nums[0]
    sum := 0
    for i := 0; i < len(nums); i++{
        sum = max(sum + nums[i], nums[i])
        res = max(res, sum)
    }
    return res
}

func max(a, b int)int{
    if a < b{
        return b
    }
    return a
}