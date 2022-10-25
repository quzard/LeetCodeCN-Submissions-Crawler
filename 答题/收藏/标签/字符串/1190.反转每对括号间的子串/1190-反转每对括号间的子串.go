// 栈 n2
func reverseParentheses1(s string) string {
    stack := [][]byte{}
    str := []byte{}
    for i := range s {
        if s[i] == '(' {
            stack = append(stack, str)
            str = []byte{}
        } else if s[i] == ')' {
            // 反转str
            for j, n := 0, len(str); j < n/2; j++ {
                str[j], str[n-1-j] = str[n-1-j], str[j]
            }
            str = append(stack[len(stack)-1], str...)
            stack = stack[:len(stack)-1]
        } else {
            str = append(str, s[i])
        }
    }
    return string(str)
}

// 预处理括号
func reverseParentheses(s string) string {
    n := len(s)
    pair := make([]int, n)
    stack := []int{}
    for i, b := range s {
        if b == '(' {
            stack = append(stack, i)
        } else if b == ')' {
            j := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            pair[i], pair[j] = j, i
        }
    }

    ans := []byte{}
    for i, step := 0, 1; i < n; i += step {
        if s[i] == '(' || s[i] == ')' {
            i = pair[i]
            step = -step
        } else {
            ans = append(ans, s[i])
        }
    }
    return string(ans)
}