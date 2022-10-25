func maximumProduct1(nums []int) int {
    res := math.MinInt32
    num1 := nums[0] * nums[1] // ��
    num2 := nums[0] * nums[1] // ����
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
	// ���������ȡ��3�������ܳ���3�������1��ȫ����(���������)��2��ȫ����(���������)��3������һ��(������С������һ���������)
	// 1��ȫ������ȡ���������
	// 2��ȫ������ȡ����������ó��ĸ�����С
	// 3������һ����exp: [-3,-2,-1,1,2]��[-5,-4,-3,-2,-1,1]��ȡ��С�������������һ������

	// ����1��
	//return sortNums(nums)

	// ����2��
	return minMaxNums(nums)
}

// ����һ����������
func sortNums(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	l := len(nums)
	return int(math.Max(float64(nums[0] * nums[1] * nums[l - 1]), float64(nums[l - 3] * nums[l - 2] * nums[l - 1])))
}

// �����������������ҳ���������������С�����������ٰ����ֵ������������
func minMaxNums(nums []int) int {
	// min1 < min2
	min1, min2 := 1000, 1000
	// max1 > max2 > max3
	max1, max2, max3 := -1000, -1000, -1000

	for _, n := range nums {
		// �������ֵ
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

		// ������Сֵ
		if n < min1 {
			min2 = min1
			min1 = n
		} else if n < min2 {
			min2 = n
		}
	}

	return int(math.Max(float64(min1 * min2 * max1), float64(max3 * max2 * max1)))
}


