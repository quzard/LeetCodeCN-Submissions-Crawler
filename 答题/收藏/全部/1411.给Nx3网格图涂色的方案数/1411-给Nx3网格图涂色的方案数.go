func numOfWays(n int) int {
    mod := int(math.Pow(10, 9) + 7)
    fi0, fi1 := 6, 6
    for i := 1; i < n; i++{
        fi0, fi1 = (2 * fi0 + 2 * fi1) % mod, (2 * fi0 + 3 * fi1) % mod
    }
    return (fi0 + fi1) % mod
}