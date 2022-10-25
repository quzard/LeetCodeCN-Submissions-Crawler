func calculate(s string) int {
	var ops []byte
	var nums []int
	var tmpNum int
	for i := 0; i < len(s); {
		switch s[i] {
		case '+', '-':
			ops = append(ops, s[i])
			i++
		case ' ':
			i++
		case '*':
			tmpNum, i = getNum(s, i+1)
			nums[len(nums)-1] *= tmpNum
		case '/':
			tmpNum, i = getNum(s, i+1)
			nums[len(nums)-1] /= tmpNum
		default:
			tmpNum, i = getNum(s, i)
			nums = append(nums, tmpNum)
		}
	}
	tmpNum = nums[0]
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case '+':
			tmpNum += nums[i+1]
		default:
			tmpNum -= nums[i+1]
		}
	}
	return tmpNum
}

func getNum(s string, i int) (int, int) {
	var num int
loop:
	for ; i < len(s); i++ {
		switch s[i] {
		case '+', '-', '*', '/':
			break loop
		case ' ':
			continue loop
		default:
			num = num*10 + int(s[i]-'0')
		}
	}
	return num, i
}