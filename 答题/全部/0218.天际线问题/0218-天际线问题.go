// 扫描线 + 优先队列
type pair struct{ right, height int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].height > h[j].height }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	boundaries := make([]int, 0, n*2)
	for _, building := range buildings {
		boundaries = append(boundaries, building[0], building[1])
	}
	sort.Ints(boundaries)

	var ans [][]int
	idx := 0
	h := hp{}
	for _, boundary := range boundaries {
		for idx < n {
			building := buildings[idx]
			if building[0] > boundary {
				break
			}
            // 获取左边界在boundary左侧的建筑的右端点和高度
			heap.Push(&h, pair{building[1], building[2]})
			idx++
		}
        // left <= h[0] < right
		for len(h) > 0 && h[0].right <= boundary {
            // 右端点在boundary左侧
			heap.Pop(&h)
		}

		maxn := 0
		if len(h) > 0 {
			maxn = h[0].height
		}
		if len(ans) == 0 || maxn != ans[len(ans)-1][1] {
			ans = append(ans, []int{boundary, maxn})
		}
	}

	return ans
}










// 扫描线简单易懂
func getSkyline1(buildings [][]int) [][]int {
	var (
		res = [][]int{}
		pq  = [][]int{}
		pre = -1
	)
	for _, b := range buildings {
		cur := [][]int{
			{b[0], -b[2]},
			{b[1], b[2]},
		}
		pq = append(pq, cur...)
	}
	arrSort(pq)
    
	heights := []int{0}
	for _, h := range pq {
		if h[1] < 0 {// 左边缘
			heights = append(heights, -h[1])
		} else { // 右边缘
			heights = remove(heights, h[1])
		}

		maxHeight := myMax(heights)

		if pre != maxHeight {
			cur := []int{h[0], maxHeight}
			res = append(res, cur)
			pre = maxHeight
		}
	}
	return res
}
func myMax(arr []int) (_max int) {
	for _, i := range arr {
		if i > _max {
			_max = i
		}
	}
	return
}
func arrSort(pq [][]int) {
	sort.Slice(pq, func(i, j int) bool {
		if pq[i][0] == pq[j][0] {
			return pq[i][1] < pq[j][1]
		} else {
			return pq[i][0] <= pq[j][0]
		}
	})
}
func remove(arr []int, tar int) []int {
	idx := -1
	for key := 0; key <= len(arr)-1; key++ {
		if arr[key] == tar {
			idx = key
            break
		}
	}
	if idx == -1 {
		return arr
	}
	arr = append(arr[:idx], arr[idx+1:]...)
	return arr
}