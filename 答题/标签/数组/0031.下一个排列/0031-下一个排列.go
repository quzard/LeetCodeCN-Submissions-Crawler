func nextPermutation(nums []int) {
    n := len(nums)
    i := n - 2
    for i >= 0 && nums[i] >= nums[i+1]{
        i--
    }
    if i >= 0 {
        j := len(nums) - 1
        for j >= i && nums[i] >= nums[j] {
            j--
        }
        fmt.Println(i, j)
        nums[i], nums[j] = nums[j], nums[i]
    }
    reverse(nums[i+1:])
}

func reverse(nums []int) {
    l, r := 0, len(nums) - 1
    for l < r {
        nums[l], nums[r] = nums[r], nums[l]
        l++
        r--
    }
}