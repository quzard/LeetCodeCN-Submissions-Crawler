func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    n := len(prices)
    // f[i][0]: ���ϳ��й�Ʊ���������
    // f[i][1]: ���ϲ����й�Ʊ�����Ҵ����䶳���е��ۼ��������
    // f[i][2]: ���ϲ����й�Ʊ�����Ҳ����䶳���е��ۼ��������
    f := make([][3]int, n)
    f[0][0] = -prices[0]
    for i := 1; i < n; i++ {
        f[i][0] = max(f[i-1][0], f[i-1][2] - prices[i])
        f[i][1] = f[i-1][0] + prices[i]
        f[i][2] = max(f[i-1][1], f[i-1][2]) 
    }
    return max(f[n-1][1], f[n-1][2])
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
