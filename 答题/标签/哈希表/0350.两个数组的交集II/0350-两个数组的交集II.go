func intersect(nums1 []int, nums2 []int) []int {
    if len(nums1) > len(nums2){
        return intersect(nums2, nums1)
    }
    hashTable := map[int]int{}
    for _, num := range nums1{
        hashTable[num]++
    }
    res := []int{}
    for _, num := range nums2{
        if hashTable[num] > 0{
            res = append(res, num)
            hashTable[num]--
        }
    }
    return res
}
// 如果给定的数组已经排好序呢？你将如何优化你的算法？
func intersect2(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    i, j := 0, 0
    res := []int{}
    for i<len(nums1) && j<len(nums2){
        if nums1[i] == nums2[j]{
            res = append(res, nums1[i])
            i++
            j++
        }else if nums1[i] > nums2[j]{
            j++
        }else if nums1[i] < nums2[j]{
            i++
        }
    }
    return res
}


// 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
// intersect1 更优

// 如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
// intersect1 更优
