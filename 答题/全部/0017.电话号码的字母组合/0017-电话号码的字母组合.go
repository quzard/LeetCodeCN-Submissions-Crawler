func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }

    words := [][]byte{
        []byte{'a', 'b', 'c'},
        []byte{'d', 'e', 'f'},
        []byte{'g', 'h', 'i'},
        []byte{'j', 'k', 'l'},
        []byte{'m', 'n', 'o'},
        []byte{'p', 'q', 'r', 's'},
        []byte{'t', 'u', 'v'},
        []byte{'w', 'x', 'y', 'z'},
    }

    res := []string{}

    var dfs func(digits string, strs []byte)

    dfs = func(digits string, strs []byte) {
        word := words[int(digits[0]-'2')]
        for j := 0; j < len(word); j++ {
            strsNew := append(append([]byte{}, strs...), word[j])
            if len(digits) > 1 {
                dfs(digits[1:], strsNew)
            } else {
                res = append(res, string(strsNew))
            }
        }
    }
    word := words[int(digits[0]-'2')]
    for j := 0; j < len(word); j++ {
        strsNew := append([]byte{}, word[j])
        if len(digits) > 1 {
            dfs(digits[1:], strsNew)
        } else {
            res = append(res, string(strsNew))
        }
    }
    return res
}