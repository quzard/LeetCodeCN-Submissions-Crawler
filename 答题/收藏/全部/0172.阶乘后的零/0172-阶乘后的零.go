func trailingZeroes1(n int )int {
    dest := 0
    temp := 5
    for n >=temp {
        dest += n/temp
        temp*=5
    }
    return dest
}


func trailingZeroes(n int) int {
    var ret int
    for n > 0 {
        ret += n / 5
        n /= 5
    }
    return ret
}