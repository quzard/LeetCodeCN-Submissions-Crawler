func lengthOfLastWord(s string) int {
    res := 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == ' '{
            if res == 0 {
                continue
            } else {
                break
            }
        }
        res++
    }
    return res
}