func subsets(nums []int) [][]int {
    res := [][]int{[]int{}}
    l := []int{}
    var dfs func(i int)
    dfs = func(i int) {
        l = append(l, nums[i])
        res = append(res, append([]int{}, l...))
        for j := i + 1; j < len(nums); j++ {
            dfs(j)            
        }
        l = l[:len(l)-1]
    }
    for i := 0; i < len(nums); i++ {
        dfs(i)
    }
    return res
}