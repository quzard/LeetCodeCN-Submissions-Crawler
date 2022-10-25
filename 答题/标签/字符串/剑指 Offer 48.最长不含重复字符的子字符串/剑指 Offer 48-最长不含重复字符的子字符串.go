func lengthOfLongestSubstring(s string) int {
    hashMap := [128]int{}
    res := 0
    l := 0
    i := 0
    for i = 0; i < len(s); i++ {
        if hashMap[s[i]] == 0{
            hashMap[s[i]] = i + 1
        } else {
            res = max(res, i - l)
            l = max(l, hashMap[s[i]])
            hashMap[s[i]] = i + 1
        }
    }
    res = max(res, i - l)
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}