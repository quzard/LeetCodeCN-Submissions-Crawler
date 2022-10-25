import "math"

func subSort(array []int) []int {
    n := len(array)
    l, r := -1, -1
    min := math.MaxInt32
    max := math.MinInt32
    for i := n - 1; i >= 0; i-- {
        // ���ȡ���������������������±�
        if array[i] > min {
            l = i
        } else {
            min = array[i]
        }

        // ȡ��������ҵ������������±�
        num := array[n-1-i]
        if num < max {
            r = n-1-i
        } else {
            max = num
        }
    }
    return []int{l, r}
}