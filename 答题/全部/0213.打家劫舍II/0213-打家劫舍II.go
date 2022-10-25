func rob1(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp := make([][4]int, n)
	for i := 1; i < len(nums); i++ {
		dp[i][0] = dp[i-1][1] + nums[i]
		dp[i][1] = max(dp[i-1][0], dp[i-1][1])
	}
	ans := max(dp[n-1][0], dp[n-1][1])
	dp[0][0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i][0] = dp[i-1][1] + nums[i]
		dp[i][1] = max(dp[i-1][0], dp[i-1][1])
	}
	return max(ans, dp[n-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func _rob(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])
	for _, v := range nums[2:] {
		first, second = second, max(first+v, second)
	}
	return second
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(_rob(nums[:n-1]), _rob(nums[1:]))
}