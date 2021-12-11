func convert(s string, numRows int) string {
    if numRows == 1{
        return s
    }
    n := len(s)
    res := make([]byte, n)
    index := 0
    diff := numRows * 2 - 2
    for j := 0; j < numRows; j++{
        if j == 0{
            for i := 0; i < n; i += diff{
                res[index] = s[i]
                index++
            }
        }else if j == numRows - 1{
            for i := j; i < n; i += diff{
                    res[index] = s[i]
                    index++
                }
        }else{
            for i := j; i < n; i += diff{
                res[index] = s[i]
                index++
                if i + 2 * (numRows - 1 - j) < n{
                    res[index] = s[i + 2 * (numRows - 1 - j)]
                    index++
                }
            }
        }
    }
    
    return string(res)
}

func convert1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var a []uint8
	b := make([][]uint8, numRows)
	t := -1
	x := 0
	for i := 0; i < len(s); i++ {
		b[x] = append(b[x], s[i])
		if x == 0 || x == numRows-1 {
			t = -t
		}
		x += t
	}
	for i := 0; i < len(b); i++ {
		a = append(a, b[i]...)
	}
	return string(a)
}