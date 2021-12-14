func sortColors(nums []int)  {
    l, r := 0, len(nums) - 1
    for i := 0; i <= r; i++{
        for nums[i] == 0 && i > l || nums[i] == 2 && i < r{
            if nums[i] == 0 && i > l{
                nums[i], nums[l] = nums[l], nums[i]
                l++
            }
            if nums[i] == 2 && i < r{
                nums[i], nums[r] = nums[r], nums[i]
                r--
            }
        }
    }
}

func sortColors2(nums []int) {
    p0, p2 := 0, len(nums)-1
    for i := 0; i <= p2; i++ {
        for ; i <= p2 && nums[i] == 2; p2-- {
            nums[i], nums[p2] = nums[p2], nums[i]
        }
        if nums[i] == 0 {
            nums[i], nums[p0] = nums[p0], nums[i]
            p0++
        }
    }
}
