func canPartition1(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	if max > target {
		return false
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < n; i++ {
		v := nums[i]
		for j := 1; j <= target; j++ {
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][target]
}

func canPartition2(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	if max > target {
		return false
	}

	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		v := nums[i]
		for j := target; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	return dp[target]
}

// 存在子集和为sum/2 则为 ture ( sum为奇数-> false)
// 转换为 01背包 问题
// 背包容量为sum/2 物品的重量和价值都为数组的值nums[i]
// dp  if dp[sum/2] == sum/2 { return true }
// 1、dp[j] 表示前i个物品任意装入容量j的背包最大价值
// 2、dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
// 3、初始化
// 4 、遍历顺序 物品->容量从大到小

func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	W := sum / 2
	dp := make([]int, W+1)

	for i := 0; i < n; i++ {
		for j := W; j >= nums[i]; j-- { // 倒序遍历
            dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[W] == W
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}