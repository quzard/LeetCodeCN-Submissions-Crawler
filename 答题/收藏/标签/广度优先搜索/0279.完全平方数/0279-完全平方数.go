func numSquares1(n int) int {
    f := make([]int, n+1)
    for i := 1; i <= n; i++ {
        minn := math.MaxInt32
        for j := 1; j*j <= i; j++ {
            minn = min(minn, f[i-j*j])
        }
        f[i] = minn + 1
    }
    return f[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// 判断是否为完全平方数
func isPerfectSquare(x int) bool {
    y := int(math.Sqrt(float64(x)))
    return y*y == x
}

// 判断是否能表示为 4^k*(8m+7)
func checkAnswer4(x int) bool {
    for x%4 == 0 {
        x /= 4
    }
    return x%8 == 7
}

func numSquares(n int) int {
    if isPerfectSquare(n) {
        return 1
    }
    if checkAnswer4(n) {
        return 4
    }
    for i := 1; i*i <= n; i++ {
        j := n - i*i
        if isPerfectSquare(j) {
            return 2
        }
    }
    return 3
}