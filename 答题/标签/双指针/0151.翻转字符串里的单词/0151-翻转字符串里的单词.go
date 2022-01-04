func reverseWords(s string) string {
    res := ""
    for i := 0; i < len(s); i++{
        if s[i] == ' '{
            continue
        }
        word := []byte{}
        for i < len(s) && s[i] != ' '{
            word = append(word, s[i])
            i++
        }
        if res == ""{
            res = string(word)
        }else{
            res = string(word) + " " + res
        }
    }
    return res
}

func reverseWords1(s string) string {
    str := strings.Fields(strings.Trim(s, " "))
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return strings.Join(str, " ")
}