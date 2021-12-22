func getSum(a, b int) int {
    for b != 0 {
        // 得到有进位的部分，向左移一位
        carry := uint(a&b) << 1
        // 得到无进位的部分
        a ^= b
        // 进位和无进位继续相加
        b = int(carry)
    }
    return a
}