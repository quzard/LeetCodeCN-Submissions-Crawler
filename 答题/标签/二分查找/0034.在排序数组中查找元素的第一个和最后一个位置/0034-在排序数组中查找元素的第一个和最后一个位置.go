func searchRange1(nums []int, target int) []int {
    if len(nums) == 0{
        return []int{-1, -1}
    }
    
    l, r := 0, len(nums) - 1
    mid := int(l + (r - l) / 2)
    for l < r && l >=0 && r < len(nums) && nums[mid] != target{
        if nums[mid] > target{
            r = mid - 1
        }else if nums[mid] < target{
            l = mid+1
        }
        mid = int(l + (r - l) / 2)
    }
    if nums[mid] == target{
        l, r:=mid, mid
        for {
            if l >= 0 && nums[l] == target{
                l--
            }
            if r < len(nums) && nums[r] == target{
                r++
            }
            if (l < 0 || nums[l] != target) && (r == len(nums) || nums[r] != target){
                break
            }
        }
        return []int{l+1,r-1}
    }else{
        return []int{-1, -1}
    }
}




func searchRange(nums []int, target int) []int {
    return []int{searchFirstEqualElement(nums,target),
    searchLastEqualElement(nums,target)}
}

func searchFirstEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			//此时 已经找到与target 相等的元素 现在只需要判断边界条件向前减 
			if mid == 0 || nums[mid-1] != target{
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// ⼆分查找最后⼀个与 target 相等的元素，时间复杂度 O(logn)
func searchLastEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] != target) { // 找到最后⼀个与 target 相等的元素
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}