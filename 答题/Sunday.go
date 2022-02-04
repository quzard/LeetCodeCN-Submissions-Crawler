package main

import "fmt"

// sunday算法

func strStr(s1, s2 string) (int, bool) {
    len1, len2 := len(s1), len(s2)
    if len1 < len2 {
        return 0, false
    }
    i := 0
Continue:
    for i < len1 && i+len2 <= len1 {
        fmt.Println(i, s1[i:])
        for j := 0; j < len2; j++ {
            if s1[i+j] != s2[j] {
                k := i + len2
                if k >= len1 {
                    return 0, false
                }
                for z := len2 - 1; z >= 0; z-- {
                    if s2[z] == s1[k] {
                        i += len2 - z
                        continue Continue
                    }
                }
                i = k
                continue Continue
            }
        }
        return i, true
        i++
    }
    return 0, false
}

func main() {
    fmt.Println(strStr("searching substring", "substr"))
}
