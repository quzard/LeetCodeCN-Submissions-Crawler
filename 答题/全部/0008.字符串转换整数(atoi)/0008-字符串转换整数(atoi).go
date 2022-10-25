import "math"

func myAtoi(s string) int {
	NoSign := true
	sign := 0
	num := 0
    i := 0
    for ; i < len(s) && s[i] == ' '; i++ {
    }
    s = s[i:]
	for i := 0; i < len(s); i++ {
		if NoSign && (s[i] == '-' || s[i] == '+') {
			NoSign = false
			if s[i] == '-' {
				sign = 1
			}
		} else if s[i] >= '0' && s[i] <= '9' {
			num = num * 10
			num += int(s[i] - '0')
			NoSign = false
			if num >= math.MaxInt32 && sign == 0 {
				return math.MaxInt32
			}
			if num >= math.MaxInt32+1 && sign == 1 {
				return -math.MaxInt32 - 1
			}
		} else {
			break
		}
	}

	if sign == 1 {
		num = -num
	}
	return num
}

func myAtoi1(s string) int {
	var (
		res    int         //存储最终结果
		sign   int  = 1    // 正负号，默认为正(1),当为负时=-1
		NoSign bool = true //是否已经出现正负符号了
		NoNum  bool = true //是否已经出现数字了
	)

	//1、丢弃无用的前导空格
	//(一定要在大循环之前，以免将数字中间的空格读入，如“12 34”，结果为12不是1234）
	for i, v := range s {
		if v != ' ' {
			s = s[i:]
			break
		}
	}

	for i, v := range s {
		//2、检查下一个字符为正还是负号（前提是，曾经没有出现过正负符号/*没有出现过数字*/还未到字符末尾）
		if v == '-' && NoSign && NoNum && i != len(s)-1 {
			sign = -1
			NoSign = false
			continue
		}
		if v == '+' && NoSign && NoNum && i != len(s)-1 {
			sign = 1
			NoSign = false
			continue
		}

		//3、检测是否到达下一个非数字字符或到达输入的结尾,忽略字符串的其余部分
		if v < '0' || v > '9' {
			break
		}

		//4、读入的这些数字转换为整数
		res = res*10 + int(v-'0')
		NoNum = false

		//判断是否溢出(在循环内实时判断，以免溢出int64)
		if sign*res < math.MinInt32 {
			return math.MinInt32
		} else if sign*res > math.MaxInt32 {
			return math.MaxInt32
		}

	}

	return sign * res
}