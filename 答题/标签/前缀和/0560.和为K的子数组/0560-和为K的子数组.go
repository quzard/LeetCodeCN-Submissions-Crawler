func subarraySum1(nums []int, k int) int {
    res := 0
    sums := []int{}
    for _, num := range nums{
        sums = append(sums, 0)
        for i := 0; i < len(sums); i++{
            sums[i] += num
            if sums[i] == k{
                res++
            }
        }
    }
    return res
}

func subarraySum2(nums []int, k int) int {
    n := len(nums) 
    s := make([]int, n+1) 
    for i := 1 ; i <=n; i++ {
        s[i] = s[i-1] + nums[i-1]
    }

    res := 0 
    hash := map[int]int{} 
    hash[0] = 1 
    for i:=1; i<=n; i++ {
        res += hash[s[i] - k] 
        hash[s[i]]++ 
    }
    return res  
}

func subarraySum(nums []int, k int) int {
    count, pre := 0, 0
    m := map[int]int{}
    m[0] = 1
    for i := 0; i < len(nums); i++ {
        pre += nums[i]
        if _, ok := m[pre - k]; ok {
            count += m[pre - k]
        }
        m[pre] += 1
    }
    return count
}