func isValid(s string) bool {
    symbols := []byte{}
    for i:=0; i<len(s); i++{
        switch s[i]{
            case '(': symbols = append(symbols, '(')
            case '{': symbols = append(symbols, '{')
            case '[': symbols = append(symbols, '[')

            case ')': 
                    if len(symbols) == 0 || symbols[len(symbols) - 1] != '('{
                        return false
                    }else{
                        symbols = symbols[:len(symbols)-1]
                    }
            case '}':
                    if len(symbols) == 0 || symbols[len(symbols) - 1] != '{'{
                        return false
                    }else{
                        symbols = symbols[:len(symbols)-1]
                    }
            case ']':
                    if len(symbols) == 0 || symbols[len(symbols) - 1] != '['{
                        return false
                    }else{
                        symbols = symbols[:len(symbols)-1]
                    }
        }
    }
    return len(symbols) == 0
}