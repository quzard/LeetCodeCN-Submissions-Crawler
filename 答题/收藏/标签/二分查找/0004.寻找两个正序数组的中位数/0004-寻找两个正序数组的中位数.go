func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    length := len(nums1) + len(nums2)
    if length % 2 == 0 {
        return (float64(find(nums1 , nums2 , length / 2)) + float64(find(nums1 , nums2 , length / 2 + 1)))/2
    } else {
        return float64(find(nums1 , nums2 , length / 2 + 1)) 
    }
}

func find(nums1 []int, nums2 []int, k int) int{
    l1, l2 := 0, 0
    for l1 < len(nums1) || l2 < len(nums2) {
        if l1 == len(nums1) {
            return nums2[l2+k-1]
        }
        if l2 == len(nums2) {
            return nums1[l1+k-1]
        }
        if k == 1 {
            return min(nums1[l1], nums2[l2])
        }
        m1 := min(l1 + k / 2, len(nums1))- 1
        m2 := min(l2 + k / 2, len(nums2))- 1
        if nums1[m1] < nums2[m2] {
            k -= m1 - l1 + 1
            l1 = m1 + 1
        } else {
            k -= m2 - l2 + 1
            l2 = m2 + 1
        }
    }
    return 0
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}