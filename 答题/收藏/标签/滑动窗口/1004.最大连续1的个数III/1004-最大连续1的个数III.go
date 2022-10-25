func longestOnes(nums []int, k int) (ans int) {
    var left, lsum, rsum int
    for right, v := range nums {
        rsum += 1 - v // ���б߽�0����
        for lsum < rsum-k { //0�������ࡣ�����ƶ���߽�
            lsum += 1 - nums[left] // ��߽����0���� rsum - lsum = ������0�ĸ���
            left++
        }
        ans = max(ans, right-left+1)
    }
    return
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}