func isPalindrome(s string) bool {
    l, r := 0, len(s) - 1
    for l <= r {
        for l <= r && ((s[l] >= 'a' && s[l] <= 'z') || (s[l] >= 'A' && s[l] <= 'Z') || (s[l] >= '0' && s[l] <= '9')) == false {
            l++
        }
        for l <= r && ((s[r] >= 'a' && s[r] <= 'z') || (s[r] >= 'A' && s[r] <= 'Z') || (s[r] >= '0' && s[r] <= '9')) == false {
            r--
        }
        if l > r {
            break
        } 
        s1 := s[l]
        s2 := s[r]
        if (s1 >= 'A' && s1 <= 'Z') {
            s1 = s1 - 'A' + 'a'
        }
        if (s2 >= 'A' && s2 <= 'Z') {
            s2 = s2 - 'A' + 'a'
        }
        if s1 != s2 {
            return false
        }
        l++
        r--
    }
    return true
}