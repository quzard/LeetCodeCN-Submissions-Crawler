func lengthOfLongestSubstring(s string) int {
    hashTable := make(map[byte]int)
    res := 0
    begin := 0
    for i:=0; i < len(s); i++{
        if index, ok := hashTable[s[i]]; !ok{
            res = max(res, i - begin + 1)
        }else{
            if index >= begin{
                res = max(res, i - 1 - begin + 1)
                begin = index + 1
            }else{
                res = max(res, i - begin + 1)
            }
        }
        hashTable[s[i]] = i
    }
    return res
}

func max(n1, n2 int) int{
    if n1 > n2{
        return n1
    }else{
        return n2
    }
}


func lengthOfLongestSubstring1(s string) int {
	m := make(map[byte]int)
	// 记录最大子串长度
	ans := 0
	for start, end, length := 0, 0, len(s); end < length; end++ {
		// 当前字符
		b := s[end]
		if curr, ok := m[b]; ok {
			// 如果当前字符最近一次出现的下一个位置，超过start，则滑动窗口
			if curr+1 > start {
				start = curr + 1
			}
		}
		// end在start之后，计算的长度就大于0，根据比较结果更新最大子串长度
		if end-start+1 > ans {
			ans = end - start + 1
		}
		// 更新当前字符最近一次出现的位置
		m[b] = end
	}
	return ans
}
