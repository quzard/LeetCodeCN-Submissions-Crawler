func canBeIncreasing(nums []int) bool {
    count := 0
    prev := 0
    cur := 1
    for cur < len(nums) {
        if nums[cur] <= nums[prev] {
            if count == 1 {
                return false
            }
            if cur == len(nums) - 1 {
                return true
            }
            if nums[prev] >= nums[cur + 1] {
                prev--
                if prev < 0 {
                    prev = cur
                    cur++
                }
            } else {
                cur++
            }
            count++
        } else {
            prev = cur
            cur++
        }
    }
    return true
}