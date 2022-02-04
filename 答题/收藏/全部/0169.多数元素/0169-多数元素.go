func majorityElement1(nums []int) int {
    n := len(nums)
    res := nums[0]
    hashMap := map[int]int{}
    for _, num := range nums{
        hashMap[num]++
        if hashMap[num] > n / 2{
            return num
        }
        if hashMap[res] < hashMap[num]{
            res = num
        }
    }
    return res
}

func majorityElement(nums []int) int {
    var vote, result int

    for _, num := range nums {
        if vote == 0 {
            result = num
        }

        if result == num {
            vote ++
        } else {
            vote --
        }
    }

    return result
}