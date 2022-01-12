func checkValidString1(s string) bool {
    l := 0
    star := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            l++
        } else if s[i] == ')' {
            if l > 0 {
                l--
            } else if star > 0 {
                star--
            } else {
                return false
            }
        } else {
            star++
        }
    }
    r := 0
    star2 := 0
    for i := len(s) - 1; i >= 0 ; i-- {
        if s[i] == ')' {
            r++
        } else if s[i] == '(' {
            if r > 0 {
                r--
            } else if star2 > 0 {
                star2--
            } else {
                return false
            }
        } else {
            star2++
        }
    }
    return r <= star2 && l <= star
}

// 贪心
func checkValidString2(s string) bool {
    minCount, maxCount := 0, 0
    for _, ch := range s {
        if ch == '(' {
            minCount++
            maxCount++
        } else if ch == ')' {
            minCount = max(minCount-1, 0)
            maxCount--
            if maxCount < 0 {
                return false
            }
        } else {
            minCount = max(minCount-1, 0)
            maxCount++
        }
    }
    return minCount == 0
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}


//动态规划
func checkValidString3(s string) bool {
    n := len(s)
    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
    }
    for i, ch := range s {
        if ch == '*' {
            dp[i][i] = true
        }
    }

    for i := 1; i < n; i++ {
        c1, c2 := s[i-1], s[i]
        dp[i-1][i] = (c1 == '(' || c1 == '*') && (c2 == ')' || c2 == '*')
    }

    for i := n - 3; i >= 0; i-- {
        c1 := s[i]
        for j := i + 2; j < n; j++ {
            c2 := s[j]
            if (c1 == '(' || c1 == '*') && (c2 == ')' || c2 == '*') {
                dp[i][j] = dp[i+1][j-1]
            }
            for k := i; k < j && !dp[i][j]; k++ {
                dp[i][j] = dp[i][k] && dp[k+1][j]
            }
        }
    }

    return dp[0][n-1]
}


//栈
func checkValidString(s string) bool {
    var leftStk, asteriskStk []int
    for i, ch := range s {
        if ch == '(' {
            leftStk = append(leftStk, i)
        } else if ch == '*' {
            asteriskStk = append(asteriskStk, i)
        } else if len(leftStk) > 0 {
            leftStk = leftStk[:len(leftStk)-1]
        } else if len(asteriskStk) > 0 {
            asteriskStk = asteriskStk[:len(asteriskStk)-1]
        } else {
            return false
        }
    }
    i := len(leftStk) - 1
    for j := len(asteriskStk) - 1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if leftStk[i] > asteriskStk[j] {
            return false
        }
    }
    return i < 0
}