func moveZeroes1(nums []int)  {
    zeroIndex := []int{}
    for i:=0; i<len(nums); i++{
        if nums[i] == 0{
            zeroIndex = append(zeroIndex, i)
        }else{
            if len(zeroIndex) > 0{
                index := zeroIndex[0]
                nums[index], nums[i] = nums[i], nums[index]
                zeroIndex = append(zeroIndex[1:], i)
            }
        }
    }
}

func moveZeroes(nums []int)  {
    length := len(nums)
    index := -1
    for i := 0; i<length; i++ {
        if nums[i] != 0 {
            if index == -1 {
                continue
            } else {
                nums[index] = nums[i]
                nums[i] = 0
                index++
            }
        } else  {
            if index == -1 {
                index = i
            } else {
                continue
            }
        }
    }

}