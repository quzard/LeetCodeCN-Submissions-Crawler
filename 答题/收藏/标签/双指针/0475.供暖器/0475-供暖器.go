func findRadius1(houses []int, heaters []int) int {
    sort.Ints(heaters)
    sort.Ints(houses)
    res := 0
    pre := 0
    for i := 0; i < len(houses); i++ {
        length := abs(houses[i] - heaters[pre])
        for j := pre + 1; j < len(heaters) && length > 0; j++ {
            newLength := abs(houses[i] - heaters[j])
            if newLength <= length {
                length = newLength
                pre = j
            } else {
                break
            }
        }
        res = max(res, length)
    }
    return res
}
func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}
func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}


func findRadius(houses []int, heaters []int) int {
    sort.Ints(houses)
    sort.Ints(heaters)

    min := 0
    houidx, heaidx := 0, 0

    for ; houidx < len(houses); houidx++ {
        for ; heaidx < len(heaters) && houses[houidx] > heaters[heaidx]; heaidx++ {}

        if heaidx == 0 {
            min = Max(min, Abs(houses[houidx]-heaters[heaidx]))
        } else if heaidx == len(heaters) {
            min = Max(min, Abs(houses[houidx]-heaters[heaidx-1]))
        } else {
            min = Max(min, Min(Abs(houses[houidx]-heaters[heaidx-1]), Abs(houses[houidx]-heaters[heaidx])))
        }
    }

    return min
}

func Min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

// ���һ����ů��
// houres �Ƿ��ݵ�λ��
// heaters �Ǽ�������λ��
// ���Ȱ뾶���ܸ����Աߵļ�������
// �ҵ�һ����С�ļ��Ȱ뾶��ʹ�ü������ܹ�ʵ��ȫ�渲��
// ����houres �еķ��ݾͲ��ø����ˣ�
// ȷʵ������������˵ [1,5] [1,5]����ô�м��2��3��4�Ͳ��ø����ˣ����Ҵ�ʱ�뾶Ϊ0����

// ���������
// �������ݵ�ͬʱ�����������������ǰ���ݱ���Ѿ������˵�ǰ����������ô�ͱ�������һ��������
// Ȼ��ͬʱ��Ҫǰ��������ͺ���������У�����һ���ܹ����ǵ�
// ��Ҫ�İ뾶Ϊ min(abs(hourse-lheater), abs(hourse-rheater))