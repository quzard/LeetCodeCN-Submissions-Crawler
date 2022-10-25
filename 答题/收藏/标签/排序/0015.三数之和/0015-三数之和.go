func threeSum(nums []int) [][]int {
    res := [][]int{}
    sort.Ints(nums)
    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        if nums[i] > 0 {
            break
        }
        l, r := i + 1, len(nums) - 1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            if sum == 0 {
                res = append(res, []int{nums[i], nums[l], nums[r]}) 
                l++
                r--
                for l < r && l-1 >= 0 && nums[l] == nums[l-1] {
                    l++
                }
                for l < r && r+1 < len(nums) && nums[r] == nums[r+1] {
                    r--
                }
            } else if sum < 0 {
                l++
                for l < r && l-1 >= 0 && nums[l] == nums[l-1] {
                    l++
                }
            } else if sum > 0 {
                r--
                for l < r && r+1 < len(nums) && nums[r] == nums[r+1] {
                    r--
                }
            }
            
        }
    }
    return res
}