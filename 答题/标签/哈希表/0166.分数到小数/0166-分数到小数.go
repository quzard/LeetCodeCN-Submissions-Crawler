func fractionToDecimal(numerator, denominator int) string {
    if numerator % denominator == 0 {
        return strconv.Itoa(numerator/denominator)
    }
    res := []byte{}
    if (numerator < 0) != (denominator < 0) {
        res = append(res, '-')
    }
    
    // 整数部分
    numerator = abs(numerator)
    denominator = abs(denominator)
    res = append(res, strconv.Itoa(numerator/denominator)...)
    res = append(res, '.')

    // 小数部分
    m := map[int]int{}
    remainder := numerator % denominator * 10
    for remainder > 0 && m[remainder] == 0 {
        m[remainder] = len(res)
        res = append(res, strconv.Itoa(remainder/denominator)...)
        remainder = remainder % denominator * 10
    }
    if m[remainder] > 0 {
        res = append(res[:m[remainder]], append([]byte{'('}, res[m[remainder]:]...)...)
        res = append(res, ')')
    }
    return string(res)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
