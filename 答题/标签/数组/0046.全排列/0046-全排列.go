var res [][]int

func dfs(nums []int, l []int){
    // fmt.Println(nums)
    if len(nums) == 0{
        res = append(res, l)
        return
    }
    tmp := make([]int, len(nums))
    ll :=  make([]int, len(l))
    for i:=0; i<len(nums); i++{
        copy(tmp, nums)
        copy(ll, l)
        if i == len(nums) - 1{
            dfs(tmp[:i], append(ll, nums[i]))
        }else if i == 0{
            dfs(tmp[i+1:], append(ll, nums[i]))
        }else{
            dfs(append(tmp[:i], tmp[i+1:]... ), append(ll, nums[i]))
        }
    }
}

func permute1(nums []int) [][]int {
    res = [][]int{}
    tmp := make([]int, len(nums))
    for i:=0; i<len(nums); i++{
        copy(tmp, nums)
        if i == len(nums) - 1{
            dfs(tmp[:i], []int{nums[i]})
        }else if i == 0{
            dfs(tmp[i+1:], []int{nums[i]})
        }else{
            dfs(append(tmp[:i], tmp[i+1:]... ), []int{nums[i]})
        }
    }
    return res
}



func permute(nums []int) [][]int {
    n := len(nums)
    result := make([][]int, 0, permuteCnt(n))
   
    var backtrack func(sorted int)
    backtrack = func(sorted int) {
        if sorted == n{
            // 输出
            result=append(result,append([]int{},nums...))
            return
        }
        backtrack(sorted+1)
        for i:=sorted+1;i<n;i++{
            // 交换
            nums[sorted],nums[i] = nums[i],nums[sorted]
            backtrack(sorted+1)
            // 撤销交换
            nums[sorted],nums[i] = nums[i],nums[sorted]
        }
    }
    backtrack(0)
    return result
}

func permuteCnt(n int) int {
    ans := 1
    for i:=0; i<n; i++ {
        ans = ans * (i+1)
    }
    return ans
}
