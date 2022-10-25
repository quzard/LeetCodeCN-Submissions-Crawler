func isPalindrome(x int) bool {
    if x < 0 || x != 0 && x%10 == 0 {
        return false
    }
    temp := x
    num := 0
    for temp > 0 {
        num = num*10 + temp%10
        temp /= 10
    }
    return num == x
}