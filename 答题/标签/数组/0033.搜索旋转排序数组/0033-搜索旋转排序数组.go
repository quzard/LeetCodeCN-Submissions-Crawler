func search1(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    for l <= r && l >= 0 && r < len(nums){
        mid := int(l + (r - l) / 2)
        fmt.Println(nums[l], nums[mid], nums[r])
        if target < nums[l] && target > nums[r]{
            break
        }
        if nums[l] == target{
            return l
        }
        if nums[r] == target{
            return r
        }

        if nums[mid] == target{
            return mid
        }

        if target > nums[l] && target < nums[mid]{
            r = mid - 1
        }else if target > nums[l] && target > nums[mid]{
            if nums[mid] < nums[l]{
                r = mid - 1
                l++
            }else{
                l = mid + 1
                r--
            }
        }else if target < nums[r] && target < nums[mid]{
            if nums[mid] < nums[r]{
                r = mid - 1
                l++
            }else{
                l = mid + 1
                r--
            }
            
        }else if  target < nums[r] && target > nums[mid]{
            l = mid + 1
        }
    }
    return -1 
}



func search(nums []int, target int) int {
    n := len(nums)
    l, r := 0, n-1
    for l <= r {
        m := (l + r) / 2
        if nums[m] == target {
            return m
        }
        if nums[l] <= nums[m] {
            if nums[l] <= target && target < nums[m] {
                r = m - 1
            } else {
                l = m + 1
            }
        } else {
            if nums[m] < target && target <= nums[r] {
                l = m + 1
            } else {
                r = m - 1
            }
        }
    }
    return -1
}