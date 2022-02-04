// 找最大差值
func maxProfit(prices []int) int {
    res := 0
    pre := -1
    for i := 0; i < len(prices); i++{
        if pre == -1{
            pre = prices[i]
        }else{
            if prices[i] < pre{
                pre = prices[i]
            }else if prices[i] > pre{
                res = max(prices[i] - pre, res)
            }
        }
    }
    return res
}

func max(a, b int) int{
    if a > b{
        return a
    }
    return b
}
