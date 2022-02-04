import (
	"container/heap"
	"sort"
)

func avoidFlood(rains []int) []int {
	ans := make([]int, len(rains))
	pools := map[int]int{}
	sunDays := []int{}
	for i := 0; i < len(rains); i++ {
		if rains[i] == 0 {
            sunDays = append(sunDays, i)
        }else if _, ok := pools[rains[i]]; !ok {
            ans[i] = -1
			pools[rains[i]] = i
        }else{
            ans[i] = -1
            if len(sunDays) == 0 || sunDays[len(sunDays)-1] < pools[rains[i]] {
                return []int{}
            }
            for j := 0; j < len(sunDays); j++ {
                if sunDays[j] > pools[rains[i]] {
                    ans[sunDays[j]] = rains[i]
                    sunDays = append(sunDays[:j], sunDays[j+1:]...)
                    break
                }
            }
			pools[rains[i]] = i
		}
	}
	for i := 0; i < len(sunDays); i++ {
		ans[sunDays[i]] = 1
	}
	return ans
}

func avoidFlood1(rains []int) []int {
	// 遍历到发洪水的下雨天， 从小顶堆中取下标， 看前面有无机会能抽水， 取不到元素说明寄
	h := &IS{}
	heap.Init(h)
	// key湖泊和val下标
	m := make(map[int]int)
	res := make([]int, len(rains))
	for i := range res {
		res[i] = -1
	}
	for i, rainPool := range rains {
		if rainPool == 0 {
			heap.Push(h, i)
		} else if rainPoolIndex, ok := m[rainPool]; !ok {
			m[rainPool] = i
		} else { //说明当前遍历到的要发洪水
			if h.Len() > 0 {
				// 当前发洪水的下标，只有拿到大于该下标的抽水日才能阻止
				left := []int{}
				for h.Len() > 0 && rainPoolIndex >= h.IntSlice[0] {
					left = append(left, heap.Pop(h).(int))
				}
				if h.Len() == 0 {
					return []int{}
				}
				e := heap.Pop(h).(int)
				res[e] = rainPool
				// 记得更新下最新的下雨的下标
				m[rainPool] = i
				// 推回堆里
				for i := range left {
					heap.Push(h, left[i])
				}
			} else {
				return []int{}
			}
		}
	}

	// 别忘了，有机会抽水的剩余不让不抽，随便挑个湖抽掉
	for h.Len() > 0 {
		e := heap.Pop(h).(int)
		res[e] = 1
	}
	return res
}

type IS struct {
	sort.IntSlice
}

func (is *IS) Push(x interface{}) {
	is.IntSlice = append(is.IntSlice, x.(int))
}
func (is *IS) Pop() interface{} {
	e := is.IntSlice[is.Len()-1]
	is.IntSlice = is.IntSlice[:is.Len()-1]
	return e
}