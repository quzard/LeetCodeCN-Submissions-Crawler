func minMeetingRooms1(intervals [][]int) int {
    sort.Slice(intervals, func(a, b int)bool{
        if intervals[a][0] < intervals[b][0]{
            return true
        }else if intervals[a][0] == intervals[b][0]{
            return intervals[a][1] < intervals[b][1]
        }else{
            return false
        }
        
    })
    rooms := [][]int{intervals[0]}
    for i:=1; i<len(intervals); i++{
        found := false
        for j:=0; j<len(rooms);j++{
            if intervals[i][0] >= rooms[j][1]{
                rooms[j] = intervals[i]
                found = true
                break
            }
        }
        if found == false{
            rooms = append(rooms, intervals[i])
        }
    }
    return len(rooms)
}

// 转换为上下车问题
func minMeetingRooms2(intervals [][]int) (res int) {
	meetings := make([][]int, 0)
	for _, interval := range intervals {
		meetings = append(meetings, []int{interval[0], 1}, []int{interval[1], -1})
	}

	sort.Slice(meetings, func(i, j int) bool {
		var order bool
		if meetings[i][0] != meetings[j][0] {
			order = meetings[i][0] < meetings[j][0]
		} else {
			order = meetings[i][1] < meetings[j][1]
		}
		return order
	})

	maxVal := 0
	cnt := 0
	for _, meeting := range meetings {
		cnt += meeting[1]
        // 车上最多有多少人
		maxVal = Max(maxVal, cnt)
	}
	return maxVal
}

func Max(a ...int) int {
	maxVal := math.MinInt32
	if len(a) == 0 {
		return maxVal
	}
	for _, num := range a {
		if maxVal < num {
			maxVal = num
		}
	}
	return maxVal
}

/*
优先队列（堆）

时间复杂度：O(NlogN) 对会议时间进行排序，NlogN，最坏情况下将需要N个会议室，对堆进行N次插入操作，每次插入后将最值元素放入堆顶需要logN
空间复杂度:O(N) 最坏情况下需要插入N个元素到堆中，总共需要N的空间
*/

func minMeetingRooms(intervals [][]int) (res int) {
	// 根据开始时间进行排序
	sort.Slice(intervals, func(i, j int) bool {
		// 使表达式为true的i所对应元素排前面，j排后面；否则i排后，j排前
		// 令会议开始时间更小的排前面，会议开始时间更大的排后面
		return intervals[i][0] < intervals[j][0]
	})

	// 构造最小堆
	h := &Heap{}
	heap.Init(h)

	// 一开始按照会议开始时间从小到大排好序，然后逐个遍历会议时间，如果堆为空则将会议结束时间存入最小堆
	// 每次最小堆的堆顶保留的是会议结束时间最小（也就是最早结束）的会议，如果后面的会议开始时间比最小结束时间晚，则说明堆顶的会议结束后可以把
	// 房间空出来给后面的会议，如果后面会议的开始时间比堆顶的最小结束时间都小，则说明最早结束的会议还没开完后面的会议就得先开始，则需要添加房间
	for _, interval := range intervals {
		if h.Len() < 1 {
			heap.Push(h, interval[1])
		} else {
			if h.arr[0] <= interval[0] {
				heap.Pop(h)
			}
			heap.Push(h, interval[1])
		}
	}
	return h.Len()
}

type Heap struct {
	arr       []int
	isMaxHeap bool
}

// 计算序列长度
func (h Heap) Len() int {
	return len(h.arr)
}

// h.arr[i] < h.arr[j]表示小根堆，h.arr[i] > h.arr[j]表示大根堆
func (h Heap) Less(i, j int) bool {
	if !h.isMaxHeap {
		// 使表达式为true的i所对应元素排前面，j所对应元素排后面；否则i所对应元素排后，j所对应元素排前
		return h.arr[i] < h.arr[j]
	} else {
		return h.arr[i] > h.arr[j]
	}
}

// 交换元素
func (h Heap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

// 不要直接调用h.Push和Pop方法，而是调用heap.Push()和heap.Pop()
// heap.Push()和heap.Pop()内部会调用h.Push()和h.Pop()进行元素的添加和删除
// 然后再调用up(Push)或down(Pop)对重新调整顺序
func (h *Heap) Push(x interface{}) {
	h.arr = append(h.arr, x.(int))
}

// 删除元素
func (h *Heap) Pop() interface{} {
	n := len(h.arr) - 1
	x := h.arr[n]
	h.arr = h.arr[:n]
	return x
}


// 有序化
/*
时间复杂度：O(NlogN)
空间复杂度：O(N)
 */
func minMeetingRooms3(intervals [][]int) (res int) {
	startTimings, endTimings := make([]int, len(intervals)), make([]int, len(intervals))
	for i, interval := range intervals {
		startTimings[i] = interval[0]
		endTimings[i] = interval[1]
	}
	
	// 对开始时间和结束时间分别进行排序
	sort.Ints(startTimings)
	sort.Ints(endTimings)
	
	start, end := 0, 0
	useRooms := 0
	// 一般遇到大于/小于/等于的情况，在草稿纸上分三种情况进行讨论，然后每种情况列出变量的变化
	// 当开始时间要大于结束时间，则说明有一个会议先结束然后再开始新的会议，也就是一个会议先结束然后腾出空间，之后再增加一个空间
	for start < len(intervals) {
		if startTimings[start] >= endTimings[end] {
			useRooms--
			end++
		}
		start++
		useRooms++
	}
	return useRooms
}