// 思路同 15. 三数之和
// 排序 + 双指针
func threeSumClosest(nums []int, target int) int {
    if len(nums) < 3{
        return 0
    }
    res := math.MaxInt32
    sort.Ints(nums)
    for i := 0; i < len(nums); i++{
        if i > 0 && nums[i] == nums[i - 1]{
            continue
        }
        l, r := i + 1, len(nums) - 1
        for l < r{
            sum := nums[l] + nums[r] + nums[i]
            if sum == target{
                return target
            }else if sum < target{
                if abs(target - res) > abs(target - sum){
                    res = sum
                }
                l ++
            }else {
                if abs(target - res) > abs(target - sum){
                    res = sum
                }
                r --
            }
        }
    }
    return res
}
func abs(a int) int{
    if a < 0{
        a = -a
    }
    return a
}