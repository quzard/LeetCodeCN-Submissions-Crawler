func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }

    res := make([]byte, 0, len(s))
    for i := 0; i < numRows; i++ {
        if i == 0 || i == numRows-1 {
            for j := i; j < len(s); j += 2*numRows - 2 {
                res = append(res, s[j])
            }
            continue
        }
        for j := i; j < len(s); j += 2*numRows - 2 {
            res = append(res, s[j])
            if j+2*(numRows-(i+1)) < len(s) {
                res = append(res, s[j+2*(numRows-(i+1))])
            }
        }
    }

    return string(res)
}