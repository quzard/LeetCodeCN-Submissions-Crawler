import (
	"container/heap"
	"sort"
)

type MedianFinder struct {
	queMax, queMin hp
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	minQ, maxQ := &mf.queMin, &mf.queMax
	// queMax �� queMin �ֱ��¼������λ��������С�ڵ�����λ��������

	//��ʱ num С�ڵ�����λ����������Ҫ��������ӵ�queMin �С�
	//�µ���λ����С�ڵ���ԭ������λ����������ǿ�����Ҫ�� queMin ���������ƶ��� queMax �С�
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {
		heap.Push(minQ, -num)
		if maxQ.Len()+1 < minQ.Len() {
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		//��ʱ num ������λ����������Ҫ��������ӵ� queMin �С�
		//�µ���λ�������ڵ���ԭ������λ����������ǿ�����Ҫ�� queMax ����С�����ƶ��� queMin �С�
		heap.Push(maxQ, num)
		if maxQ.Len() > minQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	minQ, maxQ := mf.queMin, mf.queMax
	// ���ۼ���ӵ���������Ϊ����ʱ��queMin �е����������� queMax ��һ������ʱ��λ��Ϊ queMin �Ķ�ͷ��
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	}
	// ���ۼ���ӵ���������Ϊż��ʱ���������ȶ����е�����������ͬ����ʱ��λ��Ϊ���ǵĶ�ͷ��ƽ��ֵ��
	return float64(maxQ.IntSlice[0]-minQ.IntSlice[0]) / 2
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */