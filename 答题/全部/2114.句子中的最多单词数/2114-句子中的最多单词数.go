func mostWordsFound2(sentences []string) int {
    res := 0
    for _, sentence := range sentences{
        res = max(res, len(strings.Split(sentence, " ")))
    }
    return res
}
func max(a, b int)int{
    if a > b{
        return a
    }
    return b
}

func mostWordsFound(sentences []string) (ans int) {
	for _, s := range sentences {
		cnt := strings.Count(s, " ") + 1
		if cnt > ans {
			ans = cnt
		}
	}
	return
}