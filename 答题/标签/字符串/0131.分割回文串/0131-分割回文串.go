func partition(s string) [][]string {
    res := [][]string{}
    dp := make([][]bool, len(s))
    for i:=0; i<len(s); i++{
        dp[i] = make([]bool, len(s))
        dp[i][i] = true
        for j:=0; j<i; j++{
            if j + 1 <= i - 1{
                dp[j][i] = (s[i] == s[j]) && dp[j + 1][i - 1]
            }else{
                dp[j][i] = (s[i] == s[j])
            }
        }
    }
    splits := []string{}
    var dfs func(index int)
    dfs = func(index int) {
        if index == len(s){
            res = append(res, append([]string{}, splits...))
        }
        for i := index; i < len(s); i++{
            if dp[index][i]{
                splits = append(splits, s[index : i + 1])
                dfs(i + 1)
                splits = splits[:len(splits) - 1]
            }
        }
    }
    dfs(0)
    return res
}



func partition1(s string) [][]string {
    n := len(s) 
    f := make([][]int8, n)
    for i := range f {
        f[i] = make([]int8, n)
    }

    var isPalindrome func(i, j int) int8
    isPalindrome = func(i, j int) int8 {
        if i >= j {
            return 1
        }

        if f[i][j] != 0 {
            return f[i][j]
        }

        f[i][j] = -1
        if s[i] == s[j] {
            f[i][j] = isPalindrome(i+1, j-1)
        }

        return f[i][j]
    }

    ans := [][]string{}
    splits := []string{}
    var dfs func(int)
    dfs = func(i int) {
        if i == n {
            ans = append(ans, append([]string(nil), splits...))
            return
        }
        for j := i; j < n; j ++ {
            if isPalindrome(i, j) > 0 {
                splits = append(splits, s[i:j+1])
                dfs(j + 1)
                splits = splits[:len(splits) - 1]
            }
        }
    }

    dfs(0)
    return ans
}