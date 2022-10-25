func decodeString(s string) string {
    stackStr := []string{}
    stackNum := []int{}
    str := ""
    num := 0
    for i := 0; i < len(s); i++ {
        if isNum(s[i]) {
            num = num * 10 + int(s[i] - '0')
            continue
        }
        if isStr(s[i]) {
            str += string(s[i:i+1])
            continue
        }
        if s[i] == '[' {
            stackStr = append(stackStr, str)
            str = ""

            stackNum = append(stackNum, num)
            num = 0

            continue
        }
        if s[i] == ']'{
            num = stackNum[len(stackNum)-1]
            stackNum = stackNum[:len(stackNum)-1]
            
            item := str
            str = stackStr[len(stackStr)-1]
            stackStr = stackStr[:len(stackStr)-1]

            for num > 0 {
                str += item
                num--
            }
        }
    }
    return str
}

func isNum(s byte) bool {
    return s >= '0' && s <= '9'
}

func isStr(s byte) bool {
    return (s >= 'a' && s <= 'z') || s >= 'A' && s <= 'Z'
}