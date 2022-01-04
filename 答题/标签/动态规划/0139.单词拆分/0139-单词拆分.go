func wordBreak(s string, wordDict []string) bool {
    dp := make([]bool, len(s))
    for i := 0; i < len(s); i++{
        if i > 0 && dp[i - 1] == false{
            continue
        }
        for j:=0; j < len(wordDict); j++{
            if len(wordDict[j]) > 0{
                for k:=0; k < len(wordDict[j]); k++{
                    if k < len(wordDict[j]) - 1{
                        if i+k == len(s) -1 || wordDict[j][k] != s[i+k]{
                            break
                        }
                    }else{
                        if i+k < len(s) && wordDict[j][k] == s[i+k]{
                            dp[i+k] = true
                            if i+k == len(s) - 1{
                                return true
                            }
                        }
                    }
                }
            }
        }
    }
    return false
}