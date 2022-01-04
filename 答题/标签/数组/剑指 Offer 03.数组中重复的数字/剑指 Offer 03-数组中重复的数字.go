func findRepeatNumber(nums []int) int {
    for i := 0; i < len(nums); i++{
        for nums[i] != i{
            if nums[nums[i]] == nums[i]{
                return nums[i]
            }
            nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
        }
    }
    return -1
}

func findRepeatNumber1(nums []int) int {
    m := make(map[int]bool)
    for _, num := range nums {
        if _, exists := m[num]; exists {
            return num
        }
        m[num] = true
    }
    return -1
}