func jump1(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = len(nums)
	}
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			dp[i+j] = min(dp[i+j], dp[i]+1)
		}
	}
	return dp[len(nums)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func jump(nums []int) int {
	ans, end, maxP := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxP = max(maxP, nums[i]+i)
		if i == end {
			end = maxP
			ans++
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}