import "strings"

func longestSubstring(s string, k int) int {
    ans := 0
    if s == "" {
        return ans
    }

    cnt := [26]int{}
    for _, ch := range s {
        cnt[ch-'a']++
    }

    var split byte
    for i, c := range cnt {
        if 0 < c && c < k {
            split = 'a' + byte(i)
            break
        }
    }
    if split == 0 {
        return len(s)
    }

    for _, subStr := range strings.Split(s, string(split)) {
        ans = max(ans, longestSubstring(subStr, k))
    }
    return ans
}

func longestSubstring1(s string, k int) int {
    ans := 0
    for t := 1; t <= 26; t++ {
        cnt := [26]int{}
        total := 0
        lessK := 0
        l := 0
        for r, ch := range s {
            ch -= 'a'
            if cnt[ch] == 0 {
                total++
                lessK++
            }
            cnt[ch]++
            if cnt[ch] == k {
                lessK--
            }

            for total > t {
                ch := s[l] - 'a'
                if cnt[ch] == k {
                    lessK++
                }
                cnt[ch]--
                if cnt[ch] == 0 {
                    total--
                    lessK--
                }
                l++
            }
            if lessK == 0 {
                ans = max(ans, r-l+1)
            }
        }
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}