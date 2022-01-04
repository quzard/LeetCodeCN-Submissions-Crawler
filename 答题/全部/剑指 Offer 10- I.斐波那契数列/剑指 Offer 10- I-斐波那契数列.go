func fib1(n int) int {
	if n < 2 {
		return n
	}
	mod := 1e9 + 7
	dp := []float64{0, 1}
	for i := 1; i < n; i++ {
		dp[0], dp[1] = dp[1], dp[0]+dp[1]
		for dp[1] >= mod {
			dp[1] -= mod
		}
	}
	return int(dp[1])
}

func fib(n int) int {
    const mod int = 1e9 + 7
    if n < 2 {
        return n
    }
    p, q, r := 0, 0, 1
    for i := 2; i <= n; i++ {
        p = q
        q = r
        r = (p + q) % mod
    }
    return r
}