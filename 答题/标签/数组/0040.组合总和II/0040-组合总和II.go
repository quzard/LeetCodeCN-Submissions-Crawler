import "sort"

func combinationSum2(candidates []int, target int) (ans [][]int) {
    sort.Ints(candidates)
    var freq [][2]int
    for _, num := range candidates {
        if freq == nil || num != freq[len(freq)-1][0] {
            freq = append(freq, [2]int{num, 1})
        } else {
            freq[len(freq)-1][1]++
        }
    }

    var l []int
    var dfs func(index, rest int)
    dfs = func(index, rest int) {
        if rest == 0 {
            ans = append(ans, append([]int{}, l...))
            return
        }
        if index == len(freq) || rest < freq[index][0] {
            return
        }

        dfs(index+1, rest)

        most := min(rest/freq[index][0], freq[index][1])
        for i := 1; i <= most; i++ {
            l = append(l, freq[index][0])
            dfs(index+1, rest-i*freq[index][0])
        }
        l = l[:len(l)-most]
    }
    dfs(0, target)
    return
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}