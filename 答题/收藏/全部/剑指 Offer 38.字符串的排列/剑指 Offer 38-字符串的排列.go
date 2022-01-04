// 回溯法
func permutation1(s string) []string {
    w := []byte(s)
    sort.Slice(w, func(a, b int) bool{
        return w[a] < w[b]
    })
    n := len(w)
    res := []string{}
    words := make([]byte, 0, n)
    selected := make([]bool, n)
    var dfs func(i int)
    dfs = func(i int){
        if i == n {
            res = append(res, string(words))
            return
        }
        for j := 0; j < len(w); j++{
            if selected[j] ==true || j > 0 && !selected[j-1] && w[j] == w[j-1]{
                continue
            }
            selected[j] = true
            words = append(words, w[j])
            dfs(i + 1)
            words = words[:len(words) - 1]
            selected[j] = false
        }
    }
    dfs(0)
    return res
}


func reverse(a []byte) {
    for i, n := 0, len(a); i < n/2; i++ {
        a[i], a[n-1-i] = a[n-1-i], a[i]
    }
}

func nextPermutation(nums []byte) bool {
    n := len(nums)
    i := n - 2
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }
    if i < 0 {
        return false
    }
    j := n - 1
    for j >= 0 && nums[i] >= nums[j] {
        j--
    }
    nums[i], nums[j] = nums[j], nums[i]
    reverse(nums[i+1:])
    return true
}

func permutation(s string) (ans []string) {
    t := []byte(s)
    sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
    for {
        ans = append(ans, string(t))
        if !nextPermutation(t) {
            break
        }
    }
    return
}