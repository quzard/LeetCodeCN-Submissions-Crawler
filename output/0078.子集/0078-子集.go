func subsets1(nums []int) [][]int {
    var dfs func(nums, l []int)
    res := [][]int{}
    dfs = func(nums []int, l []int){
        if len(nums) == 0{
            res = append(res, l)
            return
        }
        dfs(nums[1:], append([]int{}, l...))
        dfs(nums[1:], append(append([]int{}, l...) ,nums[0]))
    }
    dfs(nums, []int{})
    return res
}

func subsets(nums []int) [][]int {
    ans := make([][]int, 1, int(math.Pow(2, float64(len(nums)))) + 1)
    ans[0] = []int{}
    for _, x := range nums {
        for _, arr := range ans {
            a := make([]int, len(arr), len(arr)+1)
            copy(a, arr)
            ans = append(ans, append(a, x))
        } 
    }
    return ans
}
