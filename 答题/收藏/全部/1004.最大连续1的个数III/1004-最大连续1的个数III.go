func longestOnes(nums []int, k int) (ans int) {
    var left, lsum, rsum int
    for right, v := range nums {
        rsum += 1 - v // 到有边界0个数
        for lsum < rsum-k { //0个数过多。向右移动左边界
            lsum += 1 - nums[left] // 左边界左边0个数 rsum - lsum = 窗口内0的个数
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