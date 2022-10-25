func minRemoveToMakeValid(s string) string {
    stack := []int{}
    for i := range s {
        if s[i] == '(' {
            stack = append(stack, i)
            continue
        }
        if s[i] == ')' {
            if len(stack) > 0 && s[stack[len(stack)-1]] == '(' {
                stack = stack[:len(stack)-1]
            }else {
                stack = append(stack, i)
            }
        }
    }
    //重点
    arr := []byte(s)
    for i := len(stack) - 1; i>= 0; i-- {
        arr = append(arr[:stack[i]], arr[stack[i] + 1:]...)
    }
    return string(arr)
}

func minRemoveToMakeValid1(s string) string {
    res := []byte(s)
    l := []int{}
    rm := map[int]bool{}
    for i := 0; i < len(res); i++ {
        if res[i] == '(' {
            l = append(l, i)
        } else if res[i] == ')' {
            if len(l) == 0 {
                rm[i] = true
            } else {
                l = l[:len(l)-1]
            }
        }
    }
    for i := 0; i < len(l); i++ {
        rm[l[i]] = true
    }
    ss := []byte{}
    for i := 0; i < len(res); i++ {
        if rm[i] == false {
            ss = append(ss, res[i])
        }
    }
    return string(ss)
}