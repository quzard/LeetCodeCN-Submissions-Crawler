func maxLength(arr []string) (ans int) {
    masks := []int{}
outer:
    for _, s := range arr {
        mask := 0
        for _, ch := range s {
            ch -= 'a'
            if mask>>ch&1 > 0 { // �� mask ���� ch����˵�� s �����ظ���ĸ���޷����ɿ��н�
                continue outer
            }
            mask |= 1 << ch // �� ch ���� mask ��
        }
        masks = append(masks, mask)
    }

    var backtrack func(int, int)
    backtrack = func(pos, mask int) {
        if pos == len(masks) {
            ans = max(ans, bits.OnesCount(uint(mask)))
            return
        }
        if mask&masks[pos] == 0 { // mask �� masks[pos] �޹���Ԫ��
            backtrack(pos+1, mask|masks[pos])
        }
        backtrack(pos+1, mask)
    }
    backtrack(0, 0)
    return
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func maxLength2(arr []string) (ans int) {
    masks := []int{0} // 0 ��Ӧ�մ�
outer:
    for _, s := range arr {
        mask := 0
        for _, ch := range s {
            ch -= 'a'
            if mask>>ch&1 > 0 { // �� mask ���� ch����˵�� s �����ظ���ĸ���޷����ɿ��н�
                continue outer
            }
            mask |= 1 << ch // �� ch ���� mask ��
        }
        for _, m := range masks {
            if m&mask == 0 { // m �� mask �޹���Ԫ��
                masks = append(masks, m|mask)
                ans = max(ans, bits.OnesCount(uint(m|mask)))
            }
        }
    }
    return
}