func maxSubArray(nums []int) int {
    res := math.MinInt32
    sum := 0
    for _, num := range nums {
        if sum + num <= num{
            sum = num
        }else {
            sum += num
        }
        if res < sum {
            res = sum
        }
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}