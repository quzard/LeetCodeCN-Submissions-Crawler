func search(nums []int, target int) int {
    l, r := 0, len(nums)-1
    for l <= r {
        mid := l + (r-l)/2
        if nums[mid] == target {
            return mid
        }
        if nums[l] == target {
            return l
        }
        if nums[r] == target {
            return r
        }
        if nums[mid] >= nums[l] {
            if target < nums[mid] && target > nums[l] {
                r = mid - 1
            } else {
                l = mid + 1
            }
        } else {
            if target > nums[mid] && target < nums[r] {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
    }
    return -1
}