func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i := 0; i < len(nums); i++ {
        if pre, ok := m[target-nums[i]]; ok {
            return []int{pre, i}
        }
        m[nums[i]] = i
    }
    return []int{}
}