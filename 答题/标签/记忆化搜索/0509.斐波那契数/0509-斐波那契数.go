func fib(n int) int {
	if n <= 1 {
		return n
	}
	num1, num2 := 0, 1
	for i := 2; i <= n; i++ {
		num1, num2 = num2, num1+num2
	}
	return num2
}