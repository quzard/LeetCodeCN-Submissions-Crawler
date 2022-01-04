func lengthOfLIS1(nums []int) int {
    dp := make([]int, len(nums))
    res := 1
    for i:=0; i<len(nums);i++{
        for j:=0; j<i;j++{
            if nums[i] > nums[j]{
                dp[i] = max(dp[i], dp[j] + 1)
            }else if nums[i] == nums[j]{
                dp[i] = max(dp[i], dp[j])
            }
        }
        if dp[i] == 0{
            dp[i]++
        }
        res = max(res, dp[i])
    }
    fmt.Print(dp)
    return res
}

func max(a, b int)int{
    if a < b{
        return b
    }
    return a
}






func getMax(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
func lengthOfLIS(nums []int) int {
	//输入：nums = [10,9,2,5,3,7,101,18]
	//输出：4
	//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
	var dp []int // 存放最终的递增子序列
	for _, num := range nums {
		if len(dp) == 0 || num > dp[len(dp)-1] {
			// 如果dp为空或当前num大于dp中最大的数,直接加入目前最长递增子序列
			dp = append(dp, num)
		} else {
			// 贪心算法
			l, r := 0, len(dp)-1
			pos := 0
			for l <= r {
				mid := (l + r) >> 1
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

