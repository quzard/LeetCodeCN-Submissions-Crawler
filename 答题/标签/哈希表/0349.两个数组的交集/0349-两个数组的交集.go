import "sort"

func intersection1(nums1 []int, nums2 []int) []int {
	hashTable := map[int]bool{}
	for _, num := range nums1 {
		hashTable[num] = true
	}
	res := []int{}
	for _, num := range nums2 {
		if hashTable[num] {
			res = append(res, num)
			hashTable[num] = false
		}
	}
	return res
}

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := []int{}
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if len(res) == 0 || x > res[len(res)-1] {
				res = append(res, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return res
}