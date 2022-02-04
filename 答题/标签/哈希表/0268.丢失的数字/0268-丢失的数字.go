func missingNumber1(nums []int) int {
    n := len(nums)
    res := (1 + n) * n / 2
    for _, num := range nums{
        res -= num
    }
    return res
}

//由于上述 2n+12n+1 个整数中，丢失的数字出现了一次，
//其余的数字都出现了两次，因此对上述 2n+12n+1 个整数进行按位异或运算，
//结果即为丢失的数字。

func missingNumber(nums []int) int {
    ans := 0
    // 原数组后加0到n，所有值都异或，出现过两次为0，出现一次则留下
    for i, n := range(nums) {
        ans ^= i
        ans ^= n
    }
    return ans ^ len(nums)
}