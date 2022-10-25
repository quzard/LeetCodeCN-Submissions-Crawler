func firstMissingPositive(nums []int) int {
    for i := 0; i < len(nums); i++ {
        if nums[i] <= 0 {
            nums[i] = len(nums) + 1
        }
    }

    for i := 0; i < len(nums); i++ {
        if abs(nums[i]) <= len(nums) {
            j := abs(nums[i]) - 1
            nums[j] = -abs(nums[j])
        }
    }
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            return i + 1
        }
    } 
    return len(nums) + 1
}

func abs(a int) int {
    if a < 0 {
        a = -a
    }
    return a
}