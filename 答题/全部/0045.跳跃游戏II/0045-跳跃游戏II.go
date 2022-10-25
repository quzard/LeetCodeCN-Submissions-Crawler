func jump(nums []int) int {
    ans, r, maxr := 0, 0, 0
    for i := 0; i < len(nums)-1; i++ {
        maxr = max(maxr, nums[i]+i)
        if i == r {
            ans++
            r = maxr
        }
    }
    return ans
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}