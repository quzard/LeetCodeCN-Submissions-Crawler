
func letterCombinations1(digits string) []string {
    if len(digits) == 0{
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

    dfs = func(digits string, strs []byte){
        word := words[int(digits[0] - '2')]
        for j := 0; j < len(word); j++{
            strsNew := append(append([]byte{}, strs...), word[j])
            if len(digits) > 1{
                dfs(digits[1:], strsNew)
            }else{
                res = append(res, string(strsNew))
            }
        }
    }
    word := words[int(digits[0] - '2')]
    for j := 0; j < len(word); j++{
        strsNew := append([]byte{},  word[j])
        if len(digits) > 1{
            dfs(digits[1:], strsNew)
        }else{
            res = append(res, string(strsNew))
        }
    }
    return res
}





var buf []byte
var ans []string

func letterCombinations(digits string) []string {
	buf = buf[:0]
	ans = ans[:0]
	if len(digits) == 0 {
		return ans
	}
	dfs(&digits, 0)
	return ans
}

func dfs(digits *string, n int) {
	if n == len(*digits) {
		ans = append(ans, string(buf))
		return
	}
	for _, v := range digitMap[(*digits)[n]] {
		buf = append(buf, v)
		dfs(digits, n+1)
		buf = buf[:len(buf)-1]
	}
	return
}

var digitMap = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}