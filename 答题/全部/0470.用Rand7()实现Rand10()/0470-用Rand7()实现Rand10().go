// ��ʽ���ã�num = (randN-1)*N + randN���ɵó�ÿ��num�ĸ��ʾ���ͬ���پܾ��������ɡ�
func rand10() int {
    for {
        num := (rand7()-1)*7 + rand7()
        if num <= 40 {
            return num%10+1
        }
    }
}