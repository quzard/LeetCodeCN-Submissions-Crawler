// dp 优化
// 分类讨论
func maxPoints1(points [][]int) int64 {
    ans := 0
    m := len(points[0])
    f := make([][2]int, m)
    sufMax := make([]int, m) // 后缀最大值
    for i, row := range points {
        if i == 0 {
            for j, v := range row {
                ans = max(ans, v)
                f[j][0] = v + j
                f[j][1] = v - j
            }
        } else {
            preMax := math.MinInt32
            for j, v := range row {
                preMax = max(preMax, f[j][0])
                res := max(v-j+preMax, v+j+sufMax[j]) // 左侧和右侧的最大值即为选择 points[i][j] 时的计算结果
                ans = max(ans, res) // 直接更新答案，这样下面就不直接存储 res 了，改为存储 res + j 和 res - j
                f[j][0] = res + j
                f[j][1] = res - j
            }
        }
        // 计算完一整行 f 后，对于每个位置 j，计算其右侧的所有 f[k] - k 的最大值
        // 这可以通过倒着遍历 f 求出
        sufMax[m-1] = f[m-1][1]
        for j := m - 2; j >= 0; j-- {
            sufMax[j] = max(sufMax[j+1], f[j][1])
        }
    }
    return int64(ans)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func maxPoints(points [][]int) int64 {
    m := len(points)
    n := len(points[0])
    dp_old := make([]int,n)
    dp_new := make([]int,n)
    max := 0
    copy(dp_old,points[0])
    for i:=1;i<m;i++ {
        dp_new = make([]int,n)
        max = math.MinInt32
        // j > k
        for j:=0;j<n;j++ {
            if dp_old[j] + j > max {
                max = dp_old[j] + j
            }
            
            dp_new[j] = max + points[i][j]- j
        }
        max = math.MinInt32
        // j < k
        for j:=n-1;j>=0;j-- {
            if dp_old[j] - j > max {
                max = dp_old[j] - j
            }
            if max + points[i][j] + j > dp_new[j] {
                dp_new[j] = max + points[i][j] + j
            } 
        }
        dp_old = dp_new
    }
    ans := int64(0)
    for j:=0;j<n;j++ {
        if int64(dp_old[j]) > ans {
            ans = int64(dp_old[j])
        }        
    }
    return ans
}
