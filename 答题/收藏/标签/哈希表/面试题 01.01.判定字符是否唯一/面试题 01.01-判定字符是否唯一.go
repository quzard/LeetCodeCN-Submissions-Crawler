func isUnique(astr string) bool {
    s := 0
    for _, str := range astr{
        mask := 1 << (str - 'a')
        if mask & s > 0{
            return false
        }
        s = s | mask
    }
    return true
}