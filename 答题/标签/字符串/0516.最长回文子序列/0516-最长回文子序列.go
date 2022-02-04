func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				dp[i][j] = dp[i-1][j+1] + 2
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j+1])
			}
		}
	}

	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestPalindromeSubseq1(s string) int {
	n := len(s)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := n - 2; i >= 0; i-- {
		last := 0
		for j := i + 1; j < n; j++ {
			tmp := dp[j]
			if s[i] == s[j] {
				dp[j] = last + 2
			} else {
				if dp[j] < dp[j-1] {
					dp[j] = dp[j-1]
				}
			}
			last = tmp
		}
	}
	return dp[n-1]
}