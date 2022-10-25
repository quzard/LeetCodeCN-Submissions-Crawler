func findContentChildren(g []int, s []int) int {
    res := 0
    sort.Ints(g)
    sort.Ints(s)
    i, j := 0, 0
    for i < len(g) && j < len(s){
        if g[i] <= s[j] {
            res++
            i++
            j++
        } else {
            j++
        }
    }
    return res
}