func isMatch1(s string, p string) bool {
    var dfs func (i , j int) bool
    dfs = func (i , j int) bool{
        if i == len(s) && j != len(p){
            if j < len(p) - 1 && p[j+1] == '*'{
                return dfs(i, j + 2)
            }
            return false
        }
        if i != len(s) && j == len(p){
            return false
        }
        if i == len(s) && j == len(p){
            return true
        }
        if j < len(p) - 1 && p[j+1] == '*'{
            if p[j] == s[i] || p[j] == '.'{
                return dfs(i + 1, j + 2) || dfs(i + 1, j) || dfs(i, j + 2)
            }
            return dfs(i, j + 2)
        }
        if p[j] == s[i] || p[j] == '.'{
            return dfs(i + 1, j + 1)
        }
        return false
    }
    return dfs(0, 0)
}


func isMatch(s string, p string) bool {
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}
	dp := [2][]bool{}
    dp[0] = make([]bool, len(p)+1)
    dp[1] =  make([]bool, len(p)+1)
    dp[1][0] = true
	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				dp[1][j] = dp[1][j-2]
				if !dp[1][j] && matches(i, j-1) {
					dp[1][j] =  dp[0][j]
				}
			} else if matches(i, j) {
				dp[1][j] = dp[0][j-1]
			}
		}
        dp[1], dp[0] = make([]bool, len(p)+1), dp[1]
	}
	return dp[0][len(p)]
}