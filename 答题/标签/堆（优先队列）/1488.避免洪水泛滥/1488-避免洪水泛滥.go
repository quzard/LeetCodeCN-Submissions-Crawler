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
	// ����������ˮ�������죬 ��С������ȡ�±꣬ ��ǰ�����޻����ܳ�ˮ�� ȡ����Ԫ��˵����
	h := &IS{}
	heap.Init(h)
	// key������val�±�
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
		} else { //˵����ǰ��������Ҫ����ˮ
			if h.Len() > 0 {
				// ��ǰ����ˮ���±ֻ꣬���õ����ڸ��±�ĳ�ˮ�ղ�����ֹ
				left := []int{}
				for h.Len() > 0 && rainPoolIndex >= h.IntSlice[0] {
					left = append(left, heap.Pop(h).(int))
				}
				if h.Len() == 0 {
					return []int{}
				}
				e := heap.Pop(h).(int)
				res[e] = rainPool
				// �ǵø��������µ�������±�
				m[rainPool] = i
				// �ƻض���
				for i := range left {
					heap.Push(h, left[i])
				}
			} else {
				return []int{}
			}
		}
	}

	// �����ˣ��л����ˮ��ʣ�಻�ò��飬������������
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