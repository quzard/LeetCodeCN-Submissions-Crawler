func findNumberOfLIS1(nums []int) int {
    dp := make([][2]int, len(nums))
    maxLen := 1
    res := 0
    for i := 0; i < len(nums); i++ {
        dp[i][0], dp[i][1] = 1, 1
        for j := i - 1; j >= 0 && j + 2 >= dp[i][0]; j-- {
            if nums[i] > nums[j] {
                if dp[j][0] + 1 > dp[i][0] {
                    dp[i][0] = dp[j][0] + 1
                    dp[i][1] = dp[j][1]
                } else if dp[j][0] + 1 == dp[i][0] {
                    dp[i][1] += dp[j][1]
                }
            }
        }
        if maxLen < dp[i][0]{
            res = dp[i][1]
            maxLen = dp[i][0]
        } else if maxLen == dp[i][0] {
            res += dp[i][1]
        }
    }
    return res
}


func findNumberOfLIS(nums []int) int {
    // d[i] 数组表示所有能成为长度为 i 的最长上升子序列的末尾元素的值
    d := [][]int{}
    // cnt[i][j] 记录了以 d[i][j] 为结尾的最长上升子序列的个数。
	cnt := [][]int{}
	for _, num := range nums{
		i := sort.Search(len(d), func(i int)bool{ 
            return d[i][len(d[i])-1] >= num 
            })
		c := 1
		if i > 0{
			k := sort.Search(len(d[i-1]), func(k int)bool{ 
                return d[i-1][k] < num 
                })
			c = cnt[i-1][len(cnt[i-1])-1] - cnt[i-1][k]
		}
		if i == len(d){
			d = append(d, []int{num})
			cnt = append(cnt, []int{0, c}) // 前缀0： 方便前缀和 优化
		}else{
			d[i] = append(d[i], num)
			cnt[i] = append(cnt[i], cnt[i][len(cnt[i])-1] + c)
		}
	}
	c := cnt[len(cnt) - 1]
	return c[len(c) - 1]
}