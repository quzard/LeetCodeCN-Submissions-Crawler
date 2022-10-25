import "sort"

func combinationSum1(candidates []int, target int) [][]int {
    res := [][]int{}
    sort.Ints(candidates)
    var dfs func(l []int, cur, sum int)
    dfs = func(l []int, cur, sum int) {
        if sum == target {
            res = append(res, append([]int{}, l...))
            return
        }
        if cur >= len(candidates) {
            return
        }
        for i := cur; i < len(candidates); i++ {
            if sum+candidates[i] <= target {
                l = append(l, candidates[i])
                dfs(l, i, sum+candidates[i])
                l = l[:len(l)-1]
            } else {
                break
            }
        }
    }
    dfs([]int{}, 0, 0)
    return res
}

func combinationSum(candidates []int, target int) (ans [][]int) {
    comb := []int{}
    sort.Ints(candidates)

    var dfs func(target, idx int)
    dfs = func(target, idx int) {
        if idx == len(candidates) {
            return
        }
        if target == 0 {
            ans = append(ans, append([]int(nil), comb...))
            return
        }
        if target-candidates[0] < 0 {
            return
        }
        dfs(target, idx+1)
        if target-candidates[idx] >= 0 {
            comb = append(comb, candidates[idx])
            dfs(target-candidates[idx], idx)
            comb = comb[:len(comb)-1]
        }
    }
    dfs(target, 0)
    return ans
}