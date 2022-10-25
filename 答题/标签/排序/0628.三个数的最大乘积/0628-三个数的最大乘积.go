func maximumProduct1(nums []int) int {
    res := math.MinInt32
    num1 := nums[0] * nums[1] // 正
    num2 := nums[0] * nums[1] // 负数
    numMax := max(nums[0], nums[1])
    numMin := min(nums[0], nums[1])
    for i, num := range nums{
        if i < 2{
            continue
        }
        res = max(max(res, num1 * num), num2 * num)
        num1 = max(max(num1, numMax * num), numMin * num)
        num2 = min(min(num2, numMax * num), numMin * num)
        numMax = max(num, numMax)
        numMin = min(num, numMin)
    }
    return res
}

func max(a, b int)int{
    if a > b{
        return a
    }
    return b
}

func min(a, b int)int{
    if a < b{
        return a
    }
    return b
}


func maximumProduct(nums []int) int {
	// 数组排序后取的3个数可能出现3种情况：1、全正数(最大三个数)，2、全负数(最大三个数)，3、两负一正(两个最小负数和一个最大正数)
	// 1、全正数：取最大三个数
	// 2、全负数：取最大三个数得出的负数最小
	// 3、两负一正：exp: [-3,-2,-1,1,2]、[-5,-4,-3,-2,-1,1]，取最小两个负数和最大一个正数

	// 方法1：
	//return sortNums(nums)

	// 方法2：
	return minMaxNums(nums)
}

// 方法一，排序后操作
func sortNums(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	l := len(nums)
	return int(math.Max(float64(nums[0] * nums[1] * nums[l - 1]), float64(nums[l - 3] * nums[l - 2] * nums[l - 1])))
}

// 方法二，遍历数组找出最大的三个数和最小的两个数，再按出现的三种情况操作
func minMaxNums(nums []int) int {
	// min1 < min2
	min1, min2 := 1000, 1000
	// max1 > max2 > max3
	max1, max2, max3 := -1000, -1000, -1000

	for _, n := range nums {
		// 查找最大值
		if n > max1 {
			max3 = max2
			max2 = max1
			max1 = n
		} else if n > max2 {
			max3 = max2
			max2 = n
		} else if n > max3 {
			max3 = n
		}

		// 查找最小值
		if n < min1 {
			min2 = min1
			min1 = n
		} else if n < min2 {
			min2 = n
		}
	}

	return int(math.Max(float64(min1 * min2 * max1), float64(max3 * max2 * max1)))
}


