func searchInsert(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    for l < r {
        mid := (l + r) / 2
        if nums[mid] < target {
            l = mid + 1
        } else if nums[mid] > target {
            r = mid - 1
        } else {
            return mid
        }
    }
    if nums[l] < target {
        return l + 1
    } else if nums[l] > target {
        return l
    }
    return l
}

func searchInsert1(nums []int, target int) int {
	nLen := len(nums)
	left, right := 0, nLen - 1
	ans := nLen
	for left <= right {
		mid := (right - left) / 2 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}