func lengthOfLIS1(nums []int) int {
    if len(nums) < 2 {
        return len(nums)
    }
    res := 0
    dp := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        dp[i] = 1
        for j := i - 1; j >= 0; j-- {
            if nums[j] < nums[i] {
                dp[i] = max(dp[i], dp[j] + 1)
            }
        }
        res = max(res, dp[i])
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func lengthOfLIS2(nums []int) int {
    l := []int{}
    for i := 0; i < len(nums); i++ {
        if len(l) == 0 || l[len(l)-1] < nums[i] {
            l = append(l, nums[i])
        }
        if l[len(l)-1] == nums[i] {
            continue
        }
        j := 0
        // 找到第一个dp[j] >= nums[i]
        for ; j < len(l) && l[j] < nums[i]; j++ {}
        l[j] = nums[i]
    }
    return len(l)
}

func lengthOfLIS(nums []int) int {
	// dp[i]表示长度为 i 的最长上升子序列的末尾元素的最小值，
	var dp []int 
	for _, num := range nums {
		if len(dp) == 0 || num > dp[len(dp)-1] {
			// 如果dp为空或当前num大于dp中最大的数,直接加入目前最长递增子序列
			dp = append(dp, num)
		} else {
			// 贪心算法 找到第一个dp[pos] >= num
			l, r := 0, len(dp)-1
			pos := 0
			for l <= r {
				mid := l + (r - l) / 2
				if dp[mid] >= num {
					pos = mid
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			dp[pos] = num
		}
	}
	return len(dp)
}