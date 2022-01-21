import (
	"math/rand"
	"sort"
	"time"
)

func findKthLargest2(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func findKthLargest1(nums []int, k int) int {
	pool := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if len(pool) < k {
			pool = append(pool, nums[i])
			sort.Ints(pool)
		} else {
			if nums[i] > pool[0] {
				pool[0] = nums[i]
				sort.Ints(pool)
			}
		}
	}
	return pool[0]
}

// 堆排序
func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		// 保证a[i] >= a[i*2+1] 左子节点 && a[i] >= a[i*2+2] 右子节点
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && l >= 0 && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && r >= 0 && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
        maxHeapify(a, largest, heapSize)
	}
}

// 快排优化
func findKthLargest3(nums []int, k int) int {
	// 快排的分组会返回元素位置
	rand.Seed(time.Now().UnixNano())
	return selectK(nums, 0, len(nums)-1, len(nums)-k)
}

func partition(nums []int, l, r int) int {
	// 随机选择基数
	k := l + rand.Intn(r-l+1)
	nums[k], nums[r] = nums[r], nums[k]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] <= nums[r] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

func selectK(nums []int, l, r, pos int) int {
	if l >= r {
		return nums[l]
	}
	q := partition(nums, l, r)
	if q == pos {
		return nums[q]
	}
	if q < pos {
		return selectK(nums, q+1, r, pos)
	} else {
		return selectK(nums, l, q-1, pos)
	}
}