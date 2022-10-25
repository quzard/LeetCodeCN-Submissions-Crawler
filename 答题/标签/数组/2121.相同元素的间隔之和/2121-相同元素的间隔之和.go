func getDistances1(arr []int) []int64 {
    res := make([]int64, len(arr))
    sum := map[int][]int64{}
    low := map[int][]int64{}
    for i := 0; i < len(arr); i++ {
        if len(sum[arr[i]]) == 0 {
            sum[arr[i]] = make([]int64, 2)
        }
        low[i] = make([]int64, 2)
        low[i][0] = sum[arr[i]][0]
        low[i][1] = sum[arr[i]][1]

        sum[arr[i]][0] += int64(i)
        sum[arr[i]][1]++
    }
    for i := 0; i < len(arr); i++ {
        num1 := sum[arr[i]][0] - low[i][0]
        cnt1 := sum[arr[i]][1] - low[i][1]
        res[i] += num1 - cnt1*int64(i)
        res[i] += low[i][1]*int64(i) - low[i][0]
    }
    return res
}

func getDistances2(arr []int) []int64 {
    pos := map[int][]int{}
    for i, v := range arr {
        pos[v] = append(pos[v], i) // 记录相同元素的位置
    }
    ans := make([]int64, len(arr))
    for _, p := range pos {
        sum := 0
        for _, i := range p {
            sum += i - p[0]
        }
        ans[p[0]] = int64(sum)              // 最左侧元素的间隔和
        for i, n := 1, len(p); i < n; i++ { // 计算下一个相同元素的间隔和
            sum -= (n - i*2) * (p[i] - p[i-1]) // 到右边的 n-i 个点的距离更近了，同时到左边 i 个点的距离更远了
            ans[p[i]] = int64(sum)
        }
    }
    return ans
}

func getDistances(a []int) (ans []int64) {
    ps := map[int][]int{}
    for i, v := range a {
        ps[v] = append(ps[v], i)
    }
    ans = make([]int64, len(a))
    for _, a := range ps {
        if len(a) == 1 {
            continue
        }
        s := 0
        for i := 1; i < len(a); i++ {
            s += a[i] - a[0]
        }
        l, r := 0, len(a)-1
        ans[a[0]] = int64(s)
        for i := 1; i < len(a); i++ {
            l++
            r--
            s += (l - 1) * (a[i] - a[i-1])
            s -= r * (a[i] - a[i-1])
            ans[a[i]] = int64(s)
        }
    }
    return
}
