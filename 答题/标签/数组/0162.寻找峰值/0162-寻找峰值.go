func findPeakElement(nums []int) int {
    found := false
    res := -1
    get := func(i int) int {
        if i == -1 || i == len(nums) {
            return math.MinInt64
        }
        return nums[i]
    }
    var dfs func(l, r int)

    dfs = func(l, r int){
        if found{
            return
        }
        if l < 0 || r >= len(nums) || l > r{
            return
        }
        mid := int(l + (r - l) / 2)
        if get(mid) > get(mid-1) && get(mid) > get(mid+1){
            found = true
            res  = mid
        }else{
            dfs(l, mid - 1)
            dfs(mid + 1, r)
        }
    }
    dfs(0, len(nums) - 1)
    return res
}



func findPeakElement1(nums []int) int {
    n := len(nums)

    // 辅助函数，输入下标 i，返回 nums[i] 的值
    // 方便处理 nums[-1] 以及 nums[n] 的边界情况
    get := func(i int) int {
        if i == -1 || i == n {
            return math.MinInt64
        }
        return nums[i]
    }

    left, right := 0, n-1
    for {
        mid := left + (right - left) / 2
        if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
            return mid
        }
        if get(mid) < get(mid+1) {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
}
