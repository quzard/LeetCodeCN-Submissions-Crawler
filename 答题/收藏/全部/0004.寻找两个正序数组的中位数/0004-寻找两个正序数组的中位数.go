func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
    size := len(nums1) + len(nums2)
    left := 0
    right := 0
    if size % 2 == 1{
        left = (size + 1) / 2
        right = left
    }else{
        left = size / 2
        right = size / 2 + 1
    }
    fmt.Println(left, right)
    count := 0
    num1 := 0
    num2 := 0
    i, j := 0, 0;
    num := 0
    for{
        if i < len(nums1) && j < len(nums2){
            if nums1[i] < nums2[j]{
                num = nums1[i]
                i++
            }else{
                num = nums2[j]
                j++
            }
        }else if i < len(nums1){
            num = nums1[i]
            i++
        }else if j < len(nums2){
            num = nums2[j]
            j++
        }
        count++
        if count == left{
            num1 = num
        }
        if count == right{
            num2 = num
        }
        if (count == left && count == right) || (count - 1 == left && count == right){
            break
        }
    }
    return (float64(num1) + float64(num2)) / 2
}

// 划分数组
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
    // 保证 len(nums1) <= len(nums2)
    if len(nums1) > len(nums2) {
        return findMedianSortedArrays(nums2, nums1)
    }
    low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
    for low <= high {
        nums1Mid = low + (high-low)>>1
        nums2Mid = k - nums1Mid

        // 使得  左边小于右边
        if nums1Mid > 0 && nums1[nums1Mid-1]>nums2[nums2Mid] {
            high = nums1Mid-1
        } else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] {
            low = nums1Mid + 1
        }else {
            break
        }
    }

    midLeft, midRight := 0, 0
    if nums1Mid == 0 {
        midLeft = nums2[nums2Mid-1]
    } else if nums2Mid == 0 {
        midLeft = nums1[nums1Mid-1]
    } else {
        midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
    }

    if (len(nums1)+len(nums2)) &1 == 1 {
        return float64(midLeft)
    }

    if nums1Mid == len(nums1) {
        midRight = nums2[nums2Mid]
    } else if nums2Mid == len(nums2) {
        midRight = nums1[nums1Mid]
    } else {
        midRight = min(nums1[nums1Mid], nums2[nums2Mid])
    }

    return float64(midLeft+midRight) / 2

}

func max(a, b int) int{
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b{
        return a
    }
    return b
}



// 二分法, 分治
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    totalLength := len(nums1) + len(nums2)
    if totalLength%2 == 1 {
        midIndex := totalLength/2
        return float64(getKthElement(nums1, nums2, midIndex + 1))
    } else {
        midIndex1, midIndex2 := totalLength/2 - 1, totalLength/2
        return float64(getKthElement(nums1, nums2, midIndex1 + 1) + getKthElement(nums1, nums2, midIndex2 + 1)) / 2.0
    }
    return 0
}

func getKthElement(nums1, nums2 []int, k int) int {
    index1, index2 := 0, 0
    for {
        if index1 == len(nums1) {
            return nums2[index2 + k - 1]
        }
        if index2 == len(nums2) {
            return nums1[index1 + k - 1]
        }
        if k == 1 {
            return min(nums1[index1], nums2[index2])
        }
        half := k/2
        newIndex1 := min(index1 + half, len(nums1)) - 1
        newIndex2 := min(index2 + half, len(nums2)) - 1
        pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
        if pivot1 <= pivot2 {
            k -= (newIndex1 - index1 + 1)
            index1 = newIndex1 + 1
        } else {
            k -= (newIndex2 - index2 + 1)
            index2 = newIndex2 + 1
        }
    }
    return 0
}

