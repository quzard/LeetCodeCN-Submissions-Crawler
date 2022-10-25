func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n := len(nums1) + len(nums2)
    if n % 2 == 1 {
        return findK(nums1, nums2, n / 2 + 1)
    }
    return (findK(nums1, nums2, n / 2 + 1) + findK(nums1, nums2, n / 2)) / 2
}

func findK(nums1 []int, nums2 []int, k int) float64 {
    i, j := 0, 0
    for k > 0 {
        if i == len(nums1) {
            return float64(nums2[j+k-1])
        }
        if j == len(nums2) {
            return float64(nums1[i+k-1])
        }
        if k == 1 {
            return float64(min(nums1[i], nums2[j]))
        }
        mid1 := min(len(nums1)-1, i + k / 2 -1)
        mid2 := min(len(nums2)-1, j + k / 2 -1)
        if nums1[mid1] < nums2[mid2] {
            k = k - (mid1 - i + 1)
            i = mid1 + 1
        } else {
            k = k - (mid2 - j + 1)
            j = mid2 + 1
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