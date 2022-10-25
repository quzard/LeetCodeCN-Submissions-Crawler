func generateMatrix1(n int) [][]int {
	dirs := [][]int{
		[]int{0, 1},  // ÓÒ
		[]int{1, 0},  // ÏÂ
		[]int{0, -1}, // ×ó
		[]int{-1, 0}, // ÉÏ
	}
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	cnt := 1
	dir := 0
	x, y := 0, 0
	head, end, left, right := -1, n, -1, n
	for cnt <= n*n {
		res[x][y] = cnt
		x += dirs[dir][0]
		y += dirs[dir][1]
		if x == head {
			x++
			y++
			left++
			dir++
		} else if x == end {
			x--
			y--
			right--
			dir++
		} else if y == left {
			y++
			x--
			end--
			dir++
		} else if y == right {
			y--
			x++
			head++
			dir++
		}
		if dir == 4 {
			dir = 0
		}
		cnt++
	}
	return res
}

func generateMatrix(n int) [][]int {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}

	l, r, t, b := 0, n-1, 0, n-1

	val := 1
	for val <= n*n {
		for i := l; i <= r; i++ {
			arr[t][i] = val
			val++
		}
		t++

		for i := t; i <= b; i++ {
			arr[i][r] = val
			val++
		}
		r--

		for i := r; i >= l; i-- {
			arr[b][i] = val
			val++
		}
		b--

		for i := b; i >= t; i-- {
			arr[i][l] = val
			val++
		}
		l++
	}

	return arr
}