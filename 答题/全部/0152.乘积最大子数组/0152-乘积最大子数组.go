func maxProduct(nums []int) int {
    if len(nums) == 0{
        return 0
    }
    res := nums[0]
    num1 := nums[0] // 最大
    num2 := nums[0] // 最小
    for i := 1; i < len(nums); i++{
        temp1 := num1
        temp2 := num2
        num1 = max(max(nums[i], temp1 * nums[i]), temp2 * nums[i])
        num2 = min(min(nums[i], temp2 * nums[i]), temp1 * nums[i])
        res = max(res, num1)
    }
    return res
}

func max(a, b int) int{
    if a > b{
        return a
    }
    return b
}

func min(a, b int) int{
    if a < b{
        return a
    }
    return b
}