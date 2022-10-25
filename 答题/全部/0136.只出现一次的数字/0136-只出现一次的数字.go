func singleNumber(nums []int) int {
    num := nums[0]
    for i := 1; i < len(nums); i++ {
        num = num ^ nums[i]
    }
    return num
}