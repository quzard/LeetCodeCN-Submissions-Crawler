func nextPermutation(nums []int)  {
    if len(nums) < 2 {
        return
    }
    // 找到第一个 nums[l] > nums[l-1]
    l := len(nums) - 1
    for ; l > 0 && nums[l] <= nums[l-1]; l-- {
    }
    if l == 0 {
        reverse(nums, 0, len(nums)-1)
        return
    }
    //记录l-1为l, 此时 nums[l+1:]为递减序列且 nums[l] < nums[l+1]
    l--
    // 找到第一个 nums[l] < nums[r]
    r := len(nums) - 1
    for ; r > l && nums[l] >= nums[r]; r-- {
    }
    nums[l], nums[r] = nums[r], nums[l]
    reverse(nums, l+1, len(nums) - 1)
}

func reverse(nums []int, l, r int) {
    for l < r {
        nums[l], nums[r] = nums[r], nums[l]
        l++
        r--
    }
}