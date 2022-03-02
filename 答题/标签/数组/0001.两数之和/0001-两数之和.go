func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for index, num := range nums {
        if _, ok := m[target - num]; ok {
            return []int{m[target - num], index}
        }
        m[num] = index
    }
    return []int{}
}