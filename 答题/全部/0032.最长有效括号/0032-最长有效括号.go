func longestValidParentheses(s string) int {
    res := 0
    stack := []int{-1}
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            stack = append(stack, i)
        } else if len(stack) == 0 {
            stack = append(stack, i)
        } else if stack[len(stack)-1] < 0 || s[stack[len(stack)-1]] == ')' {
            stack = append(stack, i)
        } else {
            stack = stack[:len(stack)-1]
            res = max(i - stack[len(stack)-1], res)
        }
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}