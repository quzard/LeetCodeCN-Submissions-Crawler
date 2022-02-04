func addStrings1(num1 string, num2 string) string {
    res := ""
    i := len(num1) - 1
    j := len(num2) - 1
    add := 0
    for i >= 0 || j >= 0{
        if i >= 0{
            add += int(num1[i] - '0')
            i--
        }
        if j >= 0{
            add += int(num2[j] - '0')
            j--
        }
        res = string('0' + byte(add % 10)) + res
        add = int(add / 10)
    }
    if add == 1{
        res = string('0' + byte(add % 10)) + res
    }
    return res
}

func addStrings(num1 string, num2 string) string {
    add := 0

    p1 := len(num1) - 1
    p2 := len(num2) - 1

    p := max(len(num1), len(num2))
    b := make([]byte, p)
    p--

    for p1 >= 0 || p2 >= 0 {
        v := add
        if p1 >= 0 {
            v += int(num1[p1] - '0')
            p1--
        }
        if p2 >= 0 {
            v += int(num2[p2] - '0')
            p2--
        }
        add = v / 10
        b[p] = byte(v % 10) + '0'   
        p--
    }

    if add > 0 {
        return "1" + string(b)
    }

    return string(b)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}