func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    n := len(nums1)
    res := 0
    hastTable := make(map[int]int,len(nums1)*len(nums2))
    for i:=0; i<n; i++{
        for j:=0; j<n; j++{
            hastTable[-nums1[i] - nums2[j]]++
        }
    }

    for i:=0; i<n; i++{
        for j:=0; j<n; j++{
            res += hastTable[nums3[i]+nums4[j]]
        }
    }
    return res
}