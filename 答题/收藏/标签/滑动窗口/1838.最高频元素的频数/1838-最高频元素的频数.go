func maxFrequency1(nums []int, k int) int {
    sort.Ints(nums)
    n := len(nums)
    res := 0
    var i, j, pre int
    i = n - 1
    pre = i
    j = i - 1
    for ; i >= 0 && i > res; i-- {
        for ; j >= 0 && k >= 0; j-- {
            if i == j {
                continue
            }
            if nums[i] - nums[j] <= k {
                k -= nums[i] - nums[j]
                pre = j
                res = max(res, i - j)
            } else {
                break
            }
        }
        if i > 0 && i - pre > 0{
            k += (i - pre) * (nums[i] - nums[i - 1])
        }
    }
    return res + 1
}
func max(a, b int) int {
    if a > b { 
        return a
    }
    return b
}


func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    ans := 1
    for l, r, total := 0, 1, 0; r < len(nums); r++ {
        total += (nums[r] - nums[r-1]) * (r - l)
        for total > k {
            total -= nums[r] - nums[l]
            l++
        }
        ans = max(ans, r-l+1)
    }
    return ans
}