func longestPalindrome1(s string) string {
    dp := make([][]bool, 2)
    res := ""
    dp[0] = make([]bool, len(s))
    dp[1] = make([]bool, len(s))
    for i := 0; i < len(s); i++ {
        for j := 0; j <= i; j++ {
            if i-j > 2 {
                dp[1][j] = (s[i] == s[j]) && dp[0][j+1]
            } else {
                dp[1][j] = s[i] == s[j]
            }
            if dp[1][j] && i+1-j > len(res) {
                res = s[j : i+1]
            }
        }
        dp[1], dp[0] = make([]bool, len(s)), dp[1]
    }
    return res
}

func longestPalindrome(s string) string {
    if s == "" {
        return ""
    }
    start, end := 0, 0
    for i := 0; i < len(s); i++ {
        //以i为中心
        left1, right1 := expandAroundCenter(s, i, i)
        //以i, i+1 为中心
        left2, right2 := expandAroundCenter(s, i, i+1)
        if right1-left1 > end-start {
            start, end = left1, right1
        }
        if right2-left2 > end-start {
            start, end = left2, right2
        }
    }
    return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
    for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
    }
    return left + 1, right - 1
}
