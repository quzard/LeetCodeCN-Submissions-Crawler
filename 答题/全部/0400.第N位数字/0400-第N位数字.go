func findNthDigit1(n int) int {
    if n <= 9{
        return n
    }
    cnt := 0
    digital := 0
    cur := 0
    pre := 0
    for cnt < n {
        pre = cnt
        cnt += (cur * 10 + 9 - cur) * (digital+1)
        digital++
        cur = cur * 10 + 9
    }
    n -= pre + 1
    num := n/digital + (cur + 1) / 10
    index := n%digital
    return int(strconv.Itoa(num)[index] - '0')
}

func findNthDigit(n int) int {
    if n == 0 {
        return 0
    }

    start, digital, count := 1, 1, 9
    
    for n > count {
        n -= count
        digital++
        start *= 10
        count = 9 * start * digital
    }

    k := start + (n - 1) / digital
    index := (n - 1) % digital

    numStr := strconv.Itoa(k)

    return int(numStr[index] - '0')
}