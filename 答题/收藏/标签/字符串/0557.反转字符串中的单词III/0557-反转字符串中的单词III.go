func reverseWords(s string) string {
    res := []byte(s)
    for i := 0; i < len(s); i++{
        if res[i] == ' '{
            continue
        }
        j := 0
        for j = i + 1; j < len(res) && res[j] != ' '; j++{
        }
        for l , r := i, j-1; l < r;{
            res[l], res[r] = res[r], res[l]
            l++
            r--
        }
        i = j
    }
    return string(res)
}

func reverseWords1(s string) string {
    a := strings.Split(s, " ")
    for i, v := range a {
        a[i] = reverse(v)
    }
    return strings.Join(a, " ")
}

func reverse(s string) string {
    if len(s) == 0 {
        return s
    }
    x := []byte(s)
    for low, high := 0, len(s)-1; low < high; {
        tmp := x[high]
        x[high] = x[low]
        x[low] = tmp
        low++
        high--
    }
    return string(x)
}