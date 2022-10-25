func goodDaysToRobBank1(security []int, time int) []int {
    res := []int{}
    if time == 0{
        for i := 0; i < len(security); i++{
            res = append(res, i)
        }
        return res
    }
    dp := make([][]int, len(security))
    for i := 0; i < len(security); i++{
        dp[i] = make([]int, 2)
        if i == 0{
            dp[i][1] = 1 // �����ҵݼ�
            continue
        }
        if security[i] <= security[i - 1]{
            dp[i][1] += dp[i - 1][1] + 1
        }else{
            dp[i][1] = 1
        }
    }
    
    for i := len(security) - 1; i >= 0; i--{
        if i == len(security) - 1{
            dp[i][0] = 1 // ���ҵ���ݼ�
            continue
        }
        if security[i] <= security[i + 1]{
            dp[i][0] += dp[i + 1][0] + 1
        }else{
            dp[i][0] = 1
        }
    }
    
    for i:= time; i < len(security) - time; i++{
        if dp[i - 1][1] >= time && dp[i + 1][0] >= time && security[i] <= security[i-1] && security[i] <= security[i+1]{
            res = append(res, i)
        }
    }
    return res
}


func goodDaysToRobBank(a []int, time int) (ans []int) {
	n:= len(a)
	suf := make([]int, n)
	suf[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		v := a[i]
		if v > a[i+1] {
			suf[i] = 1
		} else {
			suf[i] =suf[i+1]+1
		}
	}

	pre := 0
	for i, v := range a {
        // ����һ�α����Ϳռ�
		if i == 0 || v > a[i-1] {
			pre = 1
		} else {
			pre++
		}

		if i-time >=0 && i+time < n {
			if pre>= time+1 && suf[i] >= time+1 {
				ans = append(ans, i)
			}
		}
	}

	return
}