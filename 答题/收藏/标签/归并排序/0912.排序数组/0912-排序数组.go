func sortArray(nums []int) []int {
    var quickSort func(left, right int)
    quickSort = func(left, right int) {
        if left >= right {
            return
        }
        l, r, mid := left, right, nums[left]
        for l < r {
            for l < r && nums[r] >= mid {
                r--
            }
            for l < r && nums[l] <= mid {
                l++
            }
            nums[l], nums[r] = nums[r], nums[l]
        }
        nums[l], nums[left] = nums[left], nums[l]
        quickSort(left, l - 1)
        quickSort(l+1, right)
    }
    quickSort(0, len(nums)-1)
    return nums
}