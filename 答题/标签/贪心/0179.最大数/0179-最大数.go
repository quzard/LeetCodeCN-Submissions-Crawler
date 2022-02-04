func largestNumber1(nums []int) string {
    numsString := make([]string, len(nums))
    onlyZero := true
    for i:=0; i<len(nums); i++{
        if nums[i] != 0{
            onlyZero = false
        }
        numsString[i] = strconv.Itoa(nums[i])
    }
    if onlyZero{
        return "0"
    }
    sort.Slice(numsString, func(i, j int) bool{
        num1 := numsString[i] + numsString[j]
        num2 := numsString[j] + numsString[i]
        return num1 > num2
    })
    return strings.Join(numsString, "")
}


func largestNumber(nums []int) string {
    sort.Slice(nums, func(i, j int) bool {
        x, y := nums[i], nums[j]
        sx, sy := 10, 10
        for sx <= x {
            sx = sx * 10
        }
        for sy <= y {
            sy = sy * 10
        }
        return sy * x + y > sx * y + x
    })

    if nums[0] == 0 {
        return "0"
    }

    var ans []byte
    for _, num := range nums {
        ans = append(ans, strconv.Itoa(num)...)
    }
    return string(ans)
}
