func permute(nums []int) [][]int {
    res := [][]int{}
    l := []int{}
    var dfs func(n int)
    dfs = func(n int) {
        if n == len(nums) {
            res = append(res, append([]int{}, l...))
            return
        }
        for i := n; i < len(nums); i++ {
            nums[n], nums[i] = nums[i], nums[n]
            l = append(l, nums[n])
            dfs(n+1)
            nums[n], nums[i] = nums[i], nums[n]
            l = l[:len(l)-1]
        }
    }
    dfs(0)
    return res
}