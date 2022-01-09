func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    for i:=0; i < len(nums); i++{
        if nums[i] > 0{
            break
        }
        if i > 0 && nums[i - 1] == nums[i]{
            continue
        }
        left, right := i + 1, len(nums) - 1
        target := -nums[i]
        // 找 nums[left] + nums[right] = -nums[i]
        for left < right{
            if nums[left] + nums[right] == target {
                res = append(res, []int{nums[i], nums[left], nums[right]})
                left++
                right--
                for left < right && nums[left] == nums[left - 1]{
                    left++
                }
                for left < right && nums[right] == nums[right + 1]{
                    right--
                }
            }else if nums[left] + nums[right] < target {
                left ++
            }else {
                right --
            }
        }

    }
    return res
}