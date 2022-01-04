func compressString1(S string) string {
    res := []string{}
    for i := 0; i < len(S); i++{
        cnt := 1
        j := i + 1
        for ; j < len(S) && S[j] == S[i]; j++{

        }
        cnt = j - i
        res = append(res, string(S[i]), strconv.Itoa(cnt))
        i = j - 1
    }
    s := strings.Join(res, "")
    if len(S) <= len(s){
        return S
    }
    return s
}

func compressString(S string) string {
	ans := []byte{}
	var curr int32 = 0
	tmp := 0
	for _, s := range S {
		if curr != 0 && curr != s {
			ans = append(ans, []byte(fmt.Sprintf("%c%d", curr, tmp))...)
			tmp = 1
		} else {
			tmp++
		}
        curr = s
	}
	ans = append(ans, []byte(fmt.Sprintf("%c%d", curr, tmp))...)
	if len(ans) >= len(S) {
		return S
	}
	return string(ans)
}

// strings.Builder{}
func compressString2(S string) string {
	builder := strings.Builder{}
	pre := ' '
	count := 0
	for _, v := range S {
		if pre == ' ' {
			pre = v
			count = 1
		} else if v == pre {
			count++
		} else {
			builder.WriteString(string(pre))
			builder.WriteString(strconv.Itoa(count))
			pre = v
			count = 1
		}
	}
	builder.WriteString(string(pre))
	builder.WriteString(strconv.Itoa(count))
	if builder.Len() < len(S) {
		return builder.String()
	} else {
		return S
	}
}