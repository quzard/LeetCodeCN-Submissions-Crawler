func moveZeroes(nums []int)  {
    l := 0
    for i := 0; i < len(nums); i++ {
        for i < len(nums) && nums[i] == 0 {
            i++
        }
        if i == len(nums) {
            break
        }
        nums[l] = nums[i]
        l++
    }
    for i := l; i < len(nums); i++ { 
        nums[i] = 0
    }
}