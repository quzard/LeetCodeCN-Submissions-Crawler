func removeInvalidParentheses(s string) []string {
	lremove, rremove := 0, 0
	for _, x := range s {
		if x == '(' {
			lremove++
		} else if x == ')' && lremove == 0 {
			rremove++
		} else if x == ')' && lremove > 0 {
			lremove--
		}
	}

	out := []string{}
	var dfs func(s string, l, r, index, lremove, rremove int)
	dfs = func(s string, l, r, index, lremove, rremove int) {
		if lremove == 0 && rremove == 0 {
			if isValid(s) {
				out = append(out, s)
			}
			return
		}
		for i := index; i < len(s); i++ {
			if s[i] == '(' {
				l++
			}
			if s[i] == ')' {
				r++
			}
			if i > index && s[i] == s[i-1] {
				continue
			}
			if s[i] == '(' && lremove > 0 {
                // 移除s[i]
				dfs(s[:i]+s[i+1:], l-1, r, i, lremove-1, rremove)
			}
			if s[i] == ')' && rremove > 0 {
                // 移除s[i]
				dfs(s[:i]+s[i+1:], l, r-1, i, lremove, rremove-1)
			}
			if r > l {
				break
			}
		}
	}

	dfs(s, 0, 0, 0, lremove, rremove)
	return out
}

func isValid(s string) bool {
	count := 0
	for _, v := range s {
		if v == '(' {
			count++
		} else if v == ')' {
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}