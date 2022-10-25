func removeDuplicates(nums []int) int {
    l := 0
    for i := 0; i < len(nums); i++ {
        nums[l] = nums[i]
        l++
        j := i + 1
        for ; j < len(nums); j++ {
            if nums[j] != nums[i] {
                break
            }
        }
        i = j - 1
    }
    return l
}