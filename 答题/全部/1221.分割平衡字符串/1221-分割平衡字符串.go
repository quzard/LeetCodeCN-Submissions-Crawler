func balancedStringSplit(s string) int {
    res := 0
    l := 0
    for i := 0; i < len(s); i++ {
        if s[i] == 'R' {
            l++
        } else{
            l--
        }
        if l == 0 {
            res++
        }
    }
    return res
}