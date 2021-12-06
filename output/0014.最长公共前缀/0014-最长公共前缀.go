func longestCommonPrefix(strs []string) string {
    i:=0
    res := -1
    for{
        if i == len(strs[0]){
            break
        }
        str := strs[0][i]
        same := true
        for j:=1; j<len(strs);j++{
            if i == len(strs[j]){
                same = false
                break
            }
            if strs[j][i] != str{
                same = false
            }
        }
        if same == true{
            res++
            i++
        }else{
            break
        }
    }
    if res == -1{
        return ""
    }
    return strs[0][:res + 1]
}