func maxProfit(prices []int) int {
    profit := 0
    minPrice := 0
    for i := 0; i < len(prices); i++ {
        if i == 0 {
            minPrice = prices[i]
            continue
        }
        if prices[i] > minPrice {
            profit = max(profit, prices[i] - minPrice)
        }
        minPrice = min(minPrice, prices[i])
    }
    return profit
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}