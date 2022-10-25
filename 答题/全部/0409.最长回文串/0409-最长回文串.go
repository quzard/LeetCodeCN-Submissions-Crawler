func longestPalindrome(s string) int {
	cnts := make([]int, 26*2)
	for _, w := range s {
		if w >= 'a' && w <= 'z' {
			cnts[int(w-'a')]++
		} else {
			cnts[26+int(w-'A')]++
		}
	}
	res := 0
	add := 0
	for _, cnt := range cnts {
		if cnt%2 == 0 {
			res += cnt
		} else {
			res += cnt - 1
			add = 1
		}
	}
	return res + add
}