import "sort"

func triangleNumber(nums []int) (ans int) {
	n := len(nums)
	sort.Ints(nums)
	for i, v := range nums {
		k := i + 1
		for j := i + 1; j < n; j++ {
			for k+1 < n && v+nums[j] > nums[k+1] {
				k++
			}
			ans += max(k-j, 0)
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func triangleNumber1(nums []int) int {
	sort.Ints(nums)
	res := 0

	le := len(nums)
	for i := le - 1; i > 1; i-- {
		if nums[i] == 0 {
			break
		}
		j := 0
		k := i - 1

		for k > j {
			if nums[k]+nums[j] <= nums[i] {
				j++
				continue
			}
			res = res + k - j
			k--
		}
	}

	return res
}