func lengthOfLIS(nums []int) int {
    l := []int{}
    for i := 0; i < len(nums); i++ {
        if len(l) == 0 || nums[i] > l[len(l)-1]{
            l = append(l, nums[i])
            continue
        }
        if nums[i] == l[len(l)-1] {
            continue
        }
        if nums[i] < l[len(l)-1] {
            for j := 0; j < len(l); j++ {
                if l[j] >= nums[i] {
                    l[j] = nums[i]
                    break
                }
            }
            
        }
    }
    return len(l)
}