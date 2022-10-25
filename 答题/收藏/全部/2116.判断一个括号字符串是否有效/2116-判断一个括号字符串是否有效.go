func canBeValid1(s string, locked string) bool {
    l, r, cL, cR := 0, 0, 0, 0
    newS := []byte{}
    for i := 0; i < len(s); i++ {
        if locked[i] == '0' {
            if s[i] == '(' {
                cL++
            } else {
                cR++
            }
        }
        if s[i] == '(' {
            l++
            newS = append(newS, '(')
        } else {
            if l > r {
                r++
                newS = append(newS, ')')
            } else {
                if cR == 0 {
                    return false
                }
                newS = append(newS, '(')
                cR--
                l++
            }
        }
    }
    if l == r {
        return true
    }
    s = string(newS)
    l, r, cL, cR = 0, 0, 0, 0
    for i := len(s) - 1; i >= 0; i-- {
        if locked[i] == '0' {
            if s[i] == '(' {
                cL++
            } else {
                cR++
            }
        }
        if s[i] == ')' {
            r++
        } else {
            if r > l {
                l++
            } else {
                if cL == 0 {
                    return false
                }
                cL--
                r++
            }
        }
    }
    return l == r
}

func canBeValid2(s string, locked string) bool {
    n := len(s)
    if n&1 == 1 {
        return false
    }
    var cntZero, cntL, cntR int
    for i := 0; i < n; i++ {
        if locked[i] == '0' {
            cntZero++
        } else {
            if s[i] == '(' {
                cntL++
            } else {
                if cntL > 0 {
                    cntL--
                } else if cntZero > 0 {
                    cntZero--
                } else {
                    return false
                }
            }
        }
    }
    cntZero = 0
    for i := n - 1; i >= 0; i-- {
        if locked[i] == '0' {
            cntZero++
        } else {
            if s[i] == ')' {
                cntR++
            } else {
                if cntR > 0 {
                    cntR--
                } else if cntZero > 0 {
                    cntZero--
                } else {
                    return false
                }
            }
        }
    }
    return true
}

func canBeValid(s string, locked string) bool {
    if len(s)%2 == 1 {
        return false
    }

    sint := func(i int) int {
        if locked[i] == '0' {
            return 0
        }
        if s[i] == '(' {
            return -1
        }
        if s[i] == ')' {
            return 1
        }
        return 0
    }
    cnt, cntz := 0, 0
    for i := 0; i < len(s); i++ {
        cnt += sint(i)
        if locked[i] == '0' {
            cntz++
        }
        if cnt > cntz {
            return false
        }
    }
    cntz, cnt = 0, 0
    for i := len(s) - 1; i >= 0; i-- {
        cnt += sint(i)
        if locked[i] == '0' {
            cntz--
        }
        if cnt < cntz {
            return false
        }
    }
    return true
}