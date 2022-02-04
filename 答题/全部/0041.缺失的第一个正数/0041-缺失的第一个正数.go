func firstMissingPositive1(nums []int) int {
    n := len(nums)
    for i := 0; i < n; i++ {
        if nums[i] <= 0 {
            nums[i] = n + 1
        }
    }
    for i := 0; i < n; i++ {
        num := abs(nums[i])
        if num <= n {
            nums[num - 1] = -abs(nums[num - 1])
        }
    }
    for i := 0; i < n; i++ {
        if nums[i] > 0 {
            return i + 1
        }
    }
    return n + 1
}

func abs(x int) int {
    if x < 0 {
        fmt.Println(x)
        return -x
    }
    return x
}

func firstMissingPositive(nums []int) int {
    n := len(nums)
    for i := 0; i < n; i++ {
        for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
            nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
        }
    }
    for i := 0; i < n; i++ {
        if nums[i] != i + 1 {
            return i + 1
        }
    }
    return n + 1
}

func firstMissingPositive2(nums []int) int {
	n := len(nums)
	for i:=0;i<n;{
		if nums[i]-1 == i {
			i++
		}else if 1<= nums[i] && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1],nums[i] = nums[i],nums[nums[i]-1]
		}else{
			i++
		}
	}
	for i,num := range nums {
		if num != i + 1{
			return i+1
		}
	}
	return n+1
}