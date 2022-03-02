func maxProfit1(prices []int) int {
    if len(prices) < 2 {
        return 0
    }
    res := 0
    l := make([]int, len(prices))
    r := make([]int, len(prices))
    l[0] = prices[0]
    r[len(prices)-1] = prices[len(prices)-1]
    for i := 1; i < len(prices); i++ {
        l[i] = min(l[i-1], prices[i])
    }
    for i := len(prices) - 2; i >= 0; i-- {
        r[i] = max(r[i+1], prices[i])
    }
    for i := 0; i < len(prices); i++ {
        res = max(r[i]-l[i], res)
    }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxProfit(prices []int) int {
	res := 0
	min := prices[0]
	for _, p := range prices {
		if p < min {
			min = p
		}
		if p > min && p - min > res {
			res = p - min
		}
	}
	return res
}