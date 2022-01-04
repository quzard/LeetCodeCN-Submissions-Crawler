func compress1(chars []byte) int {
    res := 0
    for i := 0; i < len(chars);{
        j := i + 1
        for ; j < len(chars) && chars[j] == chars[i]; j++{

        }
        cnt := j - i
        if cnt > 1{
            chars[res] = chars[i]
            res++
            nums := strconv.Itoa(cnt)
            for k := 0; k < len(nums); k++{
                chars[res] = nums[k]
                res++
            }
        }else{
            chars[res] = chars[i]
            res++
        }
        i = j
    }
    return res
}

func compress(chars []byte) int {
	num := len(chars)
	curCIndex := 0
	end := 0
	for i := 1; i <= num; i++ {
		if i == num || chars[i] != chars[curCIndex] {
			chars[end] = chars[curCIndex]
			end++
			n := i - curCIndex
			if n > 1 {
				numN := 0
				for ; n != 0; n /= 10 {
					chars[end] = '0' + byte(n%10)
					end++
					numN++
				}
				s := chars[end-numN : end]

				for i := 0; i < numN/2; i++ {
					s[i], s[numN-1-i] = s[numN-1-i], s[i]
				}
			}
			curCIndex = i
		}
	}
	return end
}