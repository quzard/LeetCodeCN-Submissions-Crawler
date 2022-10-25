func minWindow(s string, t string) string {
    if len(t) == 0 || len(s) < len(t) {
        return ""
    }
    m := map[byte]int{}
    cnt := len(t)
    for i := 0; i < len(t); i++ {
        m[t[i]]++
    }
    l := 0
    res := ""
    length := len(s)
    for r := 0; r < len(s); r++ {
        m[s[r]]--
        if m[s[r]] >= 0 {
            cnt--
        }
        for l < r && m[s[l]] < 0 {
            m[s[l]]++
            l++
        }
        if cnt == 0 && r-l+1 <= length {
            res = s[l : r+1]
            length = r - l + 1
        }
    }
    return res
}