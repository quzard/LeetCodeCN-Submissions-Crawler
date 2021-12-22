func strStr1(haystack string, needle string) int {
    if len(needle) == 0{
        return 0
    }
    if len(haystack) < len(needle){
        return -1
    }
    for i:=0; i<=len(haystack)-len(needle);i++{
        for j:=0; j<len(needle);j++{
            if haystack[i+j] != needle[j]{
                break
            }
            if j == len(needle) - 1{
                return i
            }
        }
    }
    return -1
}

// kmp
func strStr(haystack, needle string) int {
    n, m := len(haystack), len(needle)
    if m == 0 {
        return 0
    }
    pi := make([]int, m)
    for i, j := 1, 0; i < m; i++ {
        for j > 0 && needle[i] != needle[j] {
            j = pi[j-1]
        }
        if needle[i] == needle[j] {
            j++
        }
        pi[i] = j
    }
    for i, j := 0, 0; i < n; i++ {
        for j > 0 && haystack[i] != needle[j] {
            j = pi[j-1]
        }
        if haystack[i] == needle[j] {
            j++
        }
        if j == m {
            return i - m + 1
        }
    }
    return -1
}
