func removeElement(nums []int, val int) int {
    r := len(nums)
    for i := 0; i < r; i++{
        for nums[i] == val && i < r{
            nums[i], nums[r - 1] = nums[r - 1], nums[i]
            r--
        }
    }
    return r
}