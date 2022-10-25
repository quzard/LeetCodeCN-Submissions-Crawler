func removeElement(nums []int, val int) int {
    r := len(nums) - 1
    for i := 0; i <= r; i++ {
        for nums[i] == val && i <= r {
            nums[i], nums[r] = nums[r], nums[i]
            r--
        }
        if i > r {
            return i
        }
    }
    return r + 1
}