func reverse(x int) int {
    num := x
    if num < 0{
        num = -num
    }
    numReverse := 0
    for num != 0{
        numReverse = numReverse * 10
        numReverse += num % 10
        num = num /10
    }
    if numReverse >= math.MaxInt32 -1 {
        return 0
    }
    if x < 0{
        numReverse = -numReverse
    }
    
    return numReverse
}