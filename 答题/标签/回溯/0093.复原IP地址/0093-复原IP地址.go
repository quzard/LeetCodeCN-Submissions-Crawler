import (
    "strconv"
    "strings"
)

func restoreIpAddresses(s string) []string {
    res := []string{}
    if len(s) < 4 {
        return res
    }
    l := []string{}
    var dfs func(i int)
    dfs = func(i int) {
        if len(l) == 4 {
            if i == len(s) {
                res = append(res, strings.Join(l, "."))
            }
            return
        }
        if i >= len(s) {
            return
        }
        l = append(l, string(s[i]))
        dfs(i + 1)
        l = l[:len(l)-1]
        if i+1 < len(s) {
            num, _ := strconv.Atoi(s[i : i+2])
            if num >= 10 {
                l = append(l, string(s[i:i+2]))
                dfs(i + 2)
                l = l[:len(l)-1]
            }
        }
        if i+2 < len(s) {
            num, _ := strconv.Atoi(s[i : i+3])
            if num >= 100 && num <= 255 {
                l = append(l, string(s[i:i+3]))
                dfs(i + 3)
                l = l[:len(l)-1]
            }
        }
    }
    dfs(0)
    return res
}