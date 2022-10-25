func twoSum1(numbers []int, target int) []int {
    l, r := 0, len(numbers) - 1
    for l < r{
        if numbers[l] + numbers[r] > target{
            r--
        }else if numbers[l] + numbers[r] < target{
            l++
        }else{
            return []int{l + 1, r + 1}
        }
        
    }
    return []int{}
}

// 二分查找 o(nlgn)
func twoSum(numbers []int, target int) []int {
    for i := 0; i < len(numbers); i++ {
        low, high := i + 1, len(numbers) - 1
        for low <= high {
            mid := (high - low) / 2 + low
            if numbers[mid] == target - numbers[i] {
                return []int{i + 1, mid + 1}
            } else if numbers[mid] > target - numbers[i] {
                high = mid - 1
            } else {
                low = mid + 1
            }
        }
    }
    return []int{-1, -1}
}