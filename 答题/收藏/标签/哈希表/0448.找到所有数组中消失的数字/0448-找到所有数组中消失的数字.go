func findDisappearedNumbers1(nums []int) []int {
    // 1 <= n <= 32
    res := []int{}
    mask := 0
    for _, num := range nums{
        if mask>>num % 2 == 1{
            continue
        }
        mask |= 1<<num 
    }
    for i := 1; i <= len(nums); i++{
        if mask>>i % 2 == 0{
            res = append(res, i)
        } 
    }
    return res
}
func findDisappearedNumbers2(nums []int) []int {
    res := []int{}
    for i := 0; i < len(nums); i++{
        if i + 1 == nums[i]{
            continue
        }
        for nums[i] != i + 1{
            if nums[i] == -9{
                break
            }
            nums[nums[i] - 1], nums[i] = nums[i], nums[nums[i] - 1]
            if nums[i] >= 1 && i != nums[i] - 1 && nums[i] == nums[nums[i] - 1]{
                nums[i] = -9
                break
            }
        }
    }
    for i := 1; i <= len(nums); i++{
        if nums[i - 1] != i{
            res = append(res, i)
        }
    }
    return res

}

func findDisappearedNumbers(nums []int) (ans []int) {
    n := len(nums)

    for _, num := range nums {
        num = (num - 1) % n // 取出原来的数
        nums[num] += n
    }
    for i, num := range nums {
        if num <= n {
            ans = append(ans, i+1)
        }
    }
    return
}
