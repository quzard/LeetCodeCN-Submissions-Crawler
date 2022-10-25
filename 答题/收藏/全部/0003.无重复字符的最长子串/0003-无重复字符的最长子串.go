func lengthOfLongestSubstring(s string) int {
    m := map[byte]int{}
    res := 0
    left := -1
    for i := 0; i < len(s); i++ {
        if index, ok := m[s[i]]; ok && index >= left {
            left = index
        }
        m[s[i]] = i
        res = max(res, i - left)
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}