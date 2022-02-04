func isVaid(s byte) bool{
    if (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z') || (s >= '0' && s <= '9'){
        return true
    }
    return false
}

func isSame(a, b byte) bool{
    if a >= '0' && a <= '9' && b >= '0' && b <= '9'{
        return a == b
    }
    if a >= 'a' && a <= 'z'  && b >= 'a' && b <= 'z'{
        return a == b
    }
    if a >= 'a' && a <= 'z' && b >= 'A' && b <= 'Z'{
        return (a - 'a') == (b - 'A')
    }
    if a >= 'A' && a <= 'Z' && b >= 'a' && b <= 'z'{
        return (a - 'A') == (b - 'a')
    }
    if a >= 'A' && a <= 'Z' && b >= 'A' && b <= 'Z'{
        return (a - 'A') == (b - 'A')
    }
    return false
}

func isPalindrome(s string) bool {
    i := 0
    j := len(s) - 1
    for i <= j{
        for i < len(s) && isVaid(s[i]) == false{
            i++
        }
        for j >= 0 && isVaid(s[j]) == false{
            j--
        }
        if i <= j{
            if isSame(s[i], s[j]) == false{
                fmt.Println(i, j, string(s[i]), string(s[j]))
                return false
            }
        }else{
            break
        }
        i++
        j--
    }
    return true
}