func plusOne1(digits []int) []int {
    if len(digits) == 0{
        return digits
    }
    i := len(digits) - 1
    if digits[i] == 9{
        for{
            if digits[i] == 9{
                if i == 0{
                    digits[i] = 0
                    digits = append(append([]int{}, 1), digits...)
                    break
                }else{
                    digits[i] = 0
                    i--
                } 
            }else{
                digits[i]++
                break
            }
        }

    }else{
        digits[i]++
    }
    return digits
}

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if carry+digits[i] == 10 {
			carry = 1
			digits[i] = 0
		} else {
			digits[i] += carry
			carry = 0
		}
	}
	if carry == 1 {
		digits[0] = 1
		digits = append(digits, 0)
	}
	return digits
}
