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
		res    int         //�洢���ս��
		sign   int  = 1    // �����ţ�Ĭ��Ϊ��(1),��Ϊ��ʱ=-1
		NoSign bool = true //�Ƿ��Ѿ���������������
		NoNum  bool = true //�Ƿ��Ѿ�����������
	)

	//1���������õ�ǰ���ո�
	//(һ��Ҫ�ڴ�ѭ��֮ǰ�����⽫�����м�Ŀո���룬�硰12 34�������Ϊ12����1234��
	for i, v := range s {
		if v != ' ' {
			s = s[i:]
			break
		}
	}

	for i, v := range s {
		//2�������һ���ַ�Ϊ�����Ǹ��ţ�ǰ���ǣ�����û�г��ֹ���������/*û�г��ֹ�����*/��δ���ַ�ĩβ��
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

		//3������Ƿ񵽴���һ���������ַ��򵽴�����Ľ�β,�����ַ��������ಿ��
		if v < '0' || v > '9' {
			break
		}

		//4���������Щ����ת��Ϊ����
		res = res*10 + int(v-'0')
		NoNum = false

		//�ж��Ƿ����(��ѭ����ʵʱ�жϣ��������int64)
		if sign*res < math.MinInt32 {
			return math.MinInt32
		} else if sign*res > math.MaxInt32 {
			return math.MaxInt32
		}

	}

	return sign * res
}