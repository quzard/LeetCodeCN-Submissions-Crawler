func movingCount1(m int, n int, k int) int {
    dp := make([][]bool, m)
    res := 0
    for i := 0; i < m; i++{
        dp[i] = make([]bool, n)
        for j := 0; j < n; j++{
            dp[i][j] = isAllowed(i, j, k)
        }
    }
    var dfs func(i, j int)
    dfs = func(i, j int){
        if i < 0 || i >= m || j < 0 || j >= n || dp[i][j] == false{
            return
        }
        res++
        dp[i][j] = false
        dfs(i + 1, j)
        dfs(i - 1, j)
        dfs(i, j - 1)
        dfs(i, j + 1)
    }
    dfs(0, 0)
    return res
}

func isAllowed(a, b, k int) bool{
    sum := 0
    for a > 0{
        sum += a % 10
        a /= 10
    }
    for b > 0{
        sum += b % 10
        b /= 10
    }
    return sum <= k
}



func movingCount(m int, n int, k int) int {
    cnt := 0
    arr := make([]bool, m*n)
    matrix := make([][]bool, m)
    for i:=0; i<m; i++ {
        matrix[i], arr = arr[:n], arr[n:]
    }

    var check func(i, j int)
    check = func(i, j int) {
        if i<0 || j<0 || i>=m || j>=n || matrix[i][j] || digitSum(i) + digitSum(j) > k{
            return
        }

        matrix[i][j] = true
        cnt++
        // 只需向右向左搜索就可以
        check(i+1, j)
        check(i, j+1)
    }
    check(0, 0)
    return cnt
}

func digitSum(n int) int {
    sum := 0
    for n > 0 {
        sum += n%10
        n /= 10
    }
    return sum
}