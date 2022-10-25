func countSegments(s string) int {
    res := 0 
    for i := 0; i < len(s); i++ {
        if s[i] != ' ' {
            res++
            for i < len(s) && s[i] != ' '{
                i++
            }
        }
    }
    return res
}