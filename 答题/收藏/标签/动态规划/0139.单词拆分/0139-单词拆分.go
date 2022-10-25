func wordBreak(s string, wordDict []string) bool {
    dp := make([]bool, len(s))
    for i := 0; i < len(s); i++ {
        if i > 0 && dp[i-1] == false {
            continue
        }
        for j := 0; j < len(wordDict); j++ {
            if len(wordDict[j]) == 0 {
                continue
            }
            for k := 0; k < len(wordDict[j]); k++ {
                if k < len(wordDict[j])-1 && i+k == len(s)-1 || wordDict[j][k] != s[i+k] {
                    break
                }
                
                if k == len(wordDict[j])-1 && i+k < len(s) && wordDict[j][k] == s[i+k] {
                    if i+k == len(s)-1 {
                        return true
                    }
                    dp[i+k] = true
                }
            }
        }
    }
    return false
}