func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}


func groupAnagrams1(strs []string) [][]string {
    res := [][]string{}
    hashTable := map[string]int{}
    for _, str := range(strs){
        s := SortString(str)
        if hashTable[s] == 0{
            res = append(res, []string{str})
            hashTable[s] = len(res)
        }else{
            res[hashTable[s] - 1] = append(res[hashTable[s] - 1], str)
        }
    }
    return res
}

func groupAnagrams(strs []string) [][]string {
    anagrams := make(map[[26]int][]string)
    for _, str := range strs {
        var chars[26]int
        for _, c := range str {
            chars[c-'a']++
        }
        anagrams[chars] = append(anagrams[chars], str)
    }
    ans := [][]string{}
    for _, val := range anagrams {
        ans = append(ans, val)
    }
    return ans
}