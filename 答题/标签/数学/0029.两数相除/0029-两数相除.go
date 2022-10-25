// ���ٳ�
// y �Ǹ�����z ������
// ���� y * z 
func quickMul(y, z int) int {
    if z == 1 {
        return y
    }
    if z == 0 {
        return 0
    }
    n := quickMul(y, z >> 1)
    if z % 2 == 0 {
        return n + n
    }
    return n + n + y
}

func divide(dividend, divisor int) int {
    if dividend == math.MinInt32 { // ���Ǳ�����Ϊ��Сֵ�����
        if divisor == 1 {
            return math.MinInt32
        }
        if divisor == -1 {
            return math.MaxInt32
        }
    }
    if divisor == math.MinInt32 { // ���ǳ���Ϊ��Сֵ�����
        if dividend == math.MinInt32 {
            return 1
        }
        return 0
    }
    if dividend == 0 { // ���Ǳ�����Ϊ 0 �����
        return 0
    }

    // һ�������ʹ�ö��ֲ���
    // �����е�����ȡ�෴����������ֻ��Ҫ����һ�����
    rev := false
    if dividend > 0 {
        dividend = -dividend
        rev = !rev
    }
    if divisor > 0 {
        divisor = -divisor
        rev = !rev
    }

    ans := 0
    left, right := 1, math.MaxInt32
    for left <= right {
        mid := left + (right-left)>>1 // ע����������Ҳ���ʹ�ó���
        if quickMul(divisor, mid) >= dividend {
            ans = mid
            if mid == math.MaxInt32 { // ע�����
                break
            }
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    if rev {
        return -ans
    }
    return ans
}