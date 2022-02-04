func longestPalindrome1(s string) string {
    dp := make([][]bool, len(s))
    for i:=0; i < len(s); i++{
        dp[i] = make([]bool, len(s))
    }
    res := ""
    for i:=len(s) - 1; i >= 0; i--{
        for j := len(s) - 1; j >=i; j--{
            if i == j{
                dp[i][i] = true
            }else{
                if i == j -1{
                    dp[i][j] = (s[i] == s[j])
                }else{
                    dp[i][j] = (s[i] == s[j]) && dp[i + 1][j - 1]
                }
            }
            if dp[i][j]{
                if len(res) < j - i + 1{
                    res = string(s[i:j+1])
                }
            }
        }
    }
    return res
}


func longestPalindrome(s string) string {
    if s == "" {
        return ""
    }
    start, end := 0, 0
    for i := 0; i < len(s); i++ {
        left1, right1 := expandAroundCenter(s, i, i)
        left2, right2 := expandAroundCenter(s, i, i + 1)
        if right1 - left1 > end - start {
            start, end = left1, right1
        }
        if right2 - left2 > end - start {
            start, end = left2, right2
        }
    }
    return s[start:end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
    for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1 , right+1 { }
    return left + 1, right - 1
}