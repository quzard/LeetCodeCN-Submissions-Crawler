func threeSum(nums []int) [][]int {
    res := [][]int{}
    if len(nums) < 3 {
        return res
    }
    sort.Ints(nums)
    for i := 0; i < len(nums) - 2; i++ {
        if nums[i] > 0 {
            return res
        }
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        l, r := i + 1, len(nums) - 1 
        for l < r {
            if nums[i] + nums[l] + nums[r] == 0 {
                res = append(res, []int{nums[i], nums[l], nums[r]})
                l++
                r--
                for l < r && nums[l] == nums[l - 1]{
                    l++
                }
                for l < r && nums[r] == nums[r + 1]{
                    r--
                }
            } else if nums[i] + nums[l] + nums[r] < 0 {
                l++
            } else {
                r--
            }
        }
    }
    return res
}