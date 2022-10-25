func minimizeTheDifference1(a [][]int, tar int) int {
	m := len(a)
    // len(DP)=最大的和=m*max(mat[i][j] )
	dp := make([]int, m*70+1)
	for i := range dp {
		dp[i] = -1e9
	}
	dp[0] = 0
	for i, row := range a {
		for j := (i + 1) * 70; j > 0; j-- {
			for _, v := range row {
				if v <= j && dp[j-v]+1 > dp[j] {
					dp[j] = dp[j-v] + 1
				}
			}
		}
	}

	ans := int(1e9)
	for i, v := range dp {
		if v == m {
			if d := abs(i - tar); d < ans {
				ans = d
			}
		}
	}
	return ans
}


func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minimizeTheDifference(mat [][]int, target int) int {
	// 集合+顺序遍历+剪枝+维护最小的大于target的值
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

    mnLarger := math.MaxInt32
	v := make([]bool, target+1)
	v[0] = true
	for _, row := range mat {
		nv := make([]bool, target+1)
		nmnLarger := math.MaxInt32
		for _, x := range row {
			nmnLarger = min(nmnLarger, mnLarger+x)
			for y := range v {
				if v[y] {
					if x+y > target {
						nmnLarger = min(nmnLarger, x+y)
					} else {
						nv[x+y] = true
					}
				}
			}
		}
		v = nv
		mnLarger = nmnLarger
	}
	res := mnLarger - target
	for x := target; x >= 0; x-- {
		if v[x] {
			return min(res, target-x)
		}
	}
	return res
}