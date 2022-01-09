func maxSubArray2(nums []int) int {
    res := nums[0]
    sum := nums[0]
    for i:=1; i<len(nums); i++{
        if sum + nums[i] <= 0 || sum + nums[i] < nums[i]{
            sum = nums[i]
        }else{
            sum += nums[i]
        }
        if res < sum{
            res = sum
        }
    }
    return res
}

func maxSubArray(nums []int) int {
    max, sum := -1 << 32, 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if max < sum {
            max = sum
        }
        if sum < 0 {
            sum = 0
        }
    }
    return max
}

func maxSubArray1(nums []int) int {
    res := nums[0]
    for i := 1; i < len(nums); i ++ {
        if nums[i] + nums[i - 1] > nums[i] {
            nums[i] += nums[i - 1]
        }
        if nums[i] > res {
            res = nums[i]
        }
    }

    return res
}