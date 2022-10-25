func maxSubsequence1(nums []int, k int) []int {
    pool := make([][]int, 0)
    for i:=0; i < len(nums); i++{
        if len(pool) < k{
            pool = append(pool, []int{nums[i], i})
            sort.Slice(pool, func(a, b int)bool{
                return pool[a][0] < pool[b][0]
            })
        }else{
            if nums[i] > pool[0][0]{
                pool[0] = []int{nums[i], i}
                sort.Slice(pool, func(a, b int)bool{
                    return pool[a][0] < pool[b][0]
                })
            }
        }
    }
    sort.Slice(pool, func(a, b int)bool{
        return pool[a][1] < pool[b][1]
    })
    res := make([]int, k)
    for i := 0; i < k; i++{
        res[i] = pool[i][0]
    }
    return res
}

func maxSubsequence(a []int, k int) (ans []int) {
	type viPair struct{ 
        v, i int 
        }
	vi := make([]viPair, len(a))

	for i, v := range a {
		vi[i] = viPair{v, i}
	}

	sort.Slice(vi, func(i, j int) bool { 
        a, b := vi[i], vi[j]; 
        return a.v > b.v || (a.v == b.v && a.i < b.i )
        })

	b := vi[:k]
	sort.Slice(b, func(i, j int) bool { 
        return b[i].i < b[j].i 
        })

	for _, p := range b {
		ans = append(ans, p.v)
	}
	return
}
