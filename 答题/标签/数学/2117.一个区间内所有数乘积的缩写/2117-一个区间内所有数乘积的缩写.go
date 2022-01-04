import (
    "fmt"
    "math"
    "math/bits"
)

func abbreviateProduct(left, right int) string {
    e, cnt2, cnt5, suf, mul := 0.0, 0, 0, 1, 1
    update := func(x int) {
        suf = suf * x % 1e5
        if mul != -1 {
            mul *= x
            if mul >= 1e10 { // 长度超过 10
                mul = -1
            }
        }
    }

    for i := left; i <= right; i++ {
        e += math.Log10(float64(i))
        x := i
        tz := bits.TrailingZeros(uint(x)) // 因子 2 的个数
        cnt2 += tz
        x >>= tz
        for ; x%5 == 0; x /= 5 {
            cnt5++ // 因子 5 的个数
        }
        update(x)
    }
    cnt10 := min(cnt2, cnt5)
    for i := cnt10; i < cnt2; i++ {
        update(2) // 补上多拆出来的 2
    }
    for i := cnt10; i < cnt5; i++ {
        update(5) // 补上多拆出来的 5
    }

    if mul != -1 { // 不需要缩写
        return fmt.Sprintf("%de%d", mul, cnt10)
    }
    pre := int(math.Pow(10, e-math.Floor(e)+4))
    return fmt.Sprintf("%d...%05de%d", pre, suf, cnt10)
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
