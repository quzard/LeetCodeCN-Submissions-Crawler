func lastRemaining1(n int, m int) int {
	if n == 1 {
		return 0
	}
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	cur := -1
	for len(arr) > 1 {
		cur = (cur + m) % len(arr)
        if cur == 0{
            arr = arr[1:]
        }else{
            arr = append(arr[:cur], arr[cur+1:]...)
        }
        cur--
	}
	return arr[0]
}

func lastRemaining(n int, m int) int {
    f := 0
    for i := 2; i <= n; i++ {
        f = (f+m)%i
    }
    return f
}