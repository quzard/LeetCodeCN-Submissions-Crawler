func executeInstructions(n int, startPos []int, s string) []int {
    res := make([]int, len(s))
    x := startPos[0]
    y := startPos[1]
    // dir := map[byte][]int{}
    // dir['L'] = []int{-1, 0}
    // dir['R'] = []int{1, 0}
    // dir['U'] = []int{0, -1}
    // dir['D'] = []int{0, 1}
    type pair struct{ x, y int }
    dir4 := []pair{
        'U': {0, -1},
        'D': {0, 1},
        'R': {1, 0},
        'L': {-1, 0},
    }
    for i := 0; i < len(s); i++ {
        x = startPos[1]
        y = startPos[0]
        p := 0
        for j := i; j < len(s); j++ {
            // x += dir[s[j]][0]
            // y += dir[s[j]][1]
            d := dir4[s[j]]
            x += d.x
            y += d.y
            if x >= 0 && x < n && y >= 0 && y < n {
                p++
            } else {
                break
            }
        }
        res[i] = p
    }
    return res
}

func executeInstructions1(n int, a []int, s string) (ans []int) {
    type pair struct{ x, y int }
    dir4 := []pair{
        'L': {0, -1},
        'R': {0, 1},
        'D': {1, 0},
        'U': {-1, 0},
    }

    ans = make([]int, len(s))
o:
    for i := range s {
        x, y := a[0], a[1]
        for j, b := range s[i:] {
            d := dir4[b]
            x += d.x
            y += d.y
            if !(0 <= x && x < n && 0 <= y && y < n) {
                ans[i] = j
                continue o
            }
        }
        ans[i] = len(s[i:])
    }
    return
}
