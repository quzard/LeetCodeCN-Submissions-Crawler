func fib(n int) int {
    if n <= 1 {
        return n
    }
    var a, b float64
    a, b = 0, 1
    mod := 1e9 + 7
    for i := 2; i <= n; i++ {
        a, b = b, a + b
        for b >= mod{
            b -= mod
        }
    }
    
    for b >= mod{
        b -= mod
    }
	return int(b)
}