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
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}
	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j-2]
				if !f[i][j] && matches(i, j-1) {
					f[i][j] =  f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}