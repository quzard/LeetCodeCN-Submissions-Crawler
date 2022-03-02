func lengthOfLongestSubstring(s string) int {
    if len(s) <= 1 {
        return len(s)
    }
    res := 0
    m := map[byte]int{}
    l := 0
    for i := 0; i < len(s); i++ {
        if _, ok := m[s[i]]; ok {
            if l <= m[s[i]] {
                l = m[s[i]] + 1
            }
        }
        res = max(res, i - l + 1)
        m[s[i]] = i
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}