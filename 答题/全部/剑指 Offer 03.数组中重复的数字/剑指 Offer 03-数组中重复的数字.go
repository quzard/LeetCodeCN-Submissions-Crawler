func findRepeatNumber(nums []int) int {
    m := map[int]bool{}
    for _, num := range nums {
        if m[num] {
            return num
        }
        m[num] = true
    }
    return 0
}