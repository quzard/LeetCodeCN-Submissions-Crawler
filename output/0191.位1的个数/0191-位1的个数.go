func hammingWeight(num uint32) int {
    res := 0
    for num > 0{
        if num % 2 == 1{
            res++
        }
        num = uint32(num / 2)
    }
    return res
}