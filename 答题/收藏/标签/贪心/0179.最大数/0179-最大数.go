func largestNumber(nums []int) string {
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