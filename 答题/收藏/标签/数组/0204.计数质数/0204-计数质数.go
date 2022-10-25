//∞£ œ…∏
func countPrimes1(n int) int {
    notPrime := make([]bool, n)
    res := 0
    for i := 2; i < n; i++{
        if !notPrime[i]{
            res++
            for j := 2 * i; j < n; j += i {
                notPrime[j] = true
            }
        }
    }
    return res
}
//∞£ œ…∏ ”≈ªØ
func countPrimes2(n int) int {
	var arr = make([]bool, n)
	var count int
	for i := 2; i*i < n; i++ {
		if !arr[i] {
			for j := i * i; j < n; j += i {
				arr[j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if !arr[i] {
			count++
		}
	}
	return count
}
//œﬂ–‘…∏
func countPrimes(n int) int {
    primes := []int{}
    notPrime := make([]bool, n)

    for i := 2; i < n; i++ {
        if !notPrime[i] {
            primes = append(primes, i)
        }
        for _, p := range primes {
            if i*p >= n {
                break
            }
            notPrime[i*p] = true
            if i%p == 0 {
                break
            }
        }
    }
    return len(primes)
}