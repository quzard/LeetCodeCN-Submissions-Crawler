func canJump1(nums []int) bool {
    if len(nums) == 0{
        return false
    }
    dp := make([]bool, len(nums))
    dp[0] = true
    for i:=0; i<len(nums);i++{
        if dp[i] == false{
            continue
        }
        num := nums[i]
        for j:=1; j+i<len(nums) && j<=num;j++{
            dp[j+i] = true
            if j+i == len(nums) - 1{
                return true
            }
        }
    }
    return dp[len(nums) - 1]
}

func canJump(nums []int) bool {
    end := 0
    for i,num := range nums{
        if i>end{
            return false
        }
        end = max(end,i+num)
    }
    return true
}

func max(i,j int)int{
    if i>j{
        return i
    }
    return j
}