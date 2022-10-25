func maxScore(nums []int) int {
	n := len(nums)
	gcb := make([][]int, n)
	for i := 0; i < n; i++ {
        gcb[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			gcb[i][j] = getGcd(nums[i], nums[j])
		}
	}

    dp := make([]int, 1 << n)
	for i := 0; i < len(dp); i++ {
		cnt := getBitOneCnt(i)
        // 如果 1 的位数不是偶数就跳过，当前的操作次数等于 cnt / 2
		if cnt % 2 == 1 {
			continue
		}
        for j := 0; j < n; j++ {
            if i & (1<<j) == 0 {
                continue
            }
            for k := j + 1; k < n; k++ {
                if i & (1 << k) == 0 {
                    continue
                }
                dp[i] = max(dp[i], dp[i - (1<<j) - (1<<k)] + gcb[j][k] * cnt / 2)
            }
        }
	}
	return dp[len(dp)-1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func getBitOneCnt(num int) int {
	res := 0
	for num > 0 {
		if num % 2 == 1 {
            res++
        }
		num /= 2
	}
	return res
}

func getGcd(x, y int) int {
	num := x % y
	if num > 0 {
		return getGcd(y, num)
	} else {
		return y
	}
}