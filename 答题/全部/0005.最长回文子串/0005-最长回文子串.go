func longestPalindrome(s string) string {
    if len(s) <= 1 {
        return s
    }
    res := ""
    for i := 0; i < len(s); i++ {
        l := Palindrome(s, i, i + 1)
        if len(l) > len(res) {
            res = l
        }
        l = Palindrome(s, i, i)
        if len(l) > len(res) {
            res = l
        }
    }
    return res
}

func Palindrome(s string, l, r int) string {
    for l >= 0 && r < len(s) && s[l] == s[r]{
        l--
        r++
    }
    return s[l+1:r]
}
