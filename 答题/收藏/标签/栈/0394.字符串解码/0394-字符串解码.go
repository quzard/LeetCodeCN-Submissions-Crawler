func decodeString(s string) string {
	numStack, strStack := make([]int, 0), make([]string, 0)
	num, str := 0, ""
	for i := 0; i < len(s); i++ {
		if isNumber(s[i]) {
			num = num*10 + int(s[i]-'0')
		} else if isLetter(s[i]) {
			str += string(s[i])
		} else if s[i] == '[' {
			numStack = append(numStack, num)
			strStack = append(strStack, str)
			num, str = 0, ""
		} else if s[i] == ']' {
			repeatTime := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			item := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			for ; repeatTime != 0; repeatTime-- {
				item += str
			}
			str = item
		}
	}
	return str
}

func isLetter(u uint8) bool {
	return 'A' <= u && u <= 'Z' || 'a' <= u && u <= 'z'
}

func isNumber(u uint8) bool {
	return '0' <= u && u <= '9'
}