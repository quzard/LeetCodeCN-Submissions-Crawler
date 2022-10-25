func sortedSquares(nums []int) []int {
    zeroIndex := 0
    for zeroIndex < len(nums) && nums[zeroIndex] < 0 {
        zeroIndex++
    }
    l, r := zeroIndex - 1, zeroIndex
    num1, num2 := 0, 0
    res := make([]int, len(nums))
    cnt := 0
    for cnt < len(nums) {
        if l >= 0 && r < len(nums){
            num1 = nums[l] * nums[l]
            num2 = nums[r] * nums[r]
            if num1 < num2 {
                l--
                res[cnt] = num1
            } else {
                r++
                res[cnt] = num2
            }
        } else if r < len(nums) {
            num2 = nums[r] * nums[r]
            r++
            res[cnt] = num2
        } else if l >= 0 {
            num1 = nums[l] * nums[l]
            l--
            res[cnt] = num1
        }
        cnt++
    }
    return res
}