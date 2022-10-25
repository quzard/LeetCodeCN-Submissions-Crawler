func isConvex(points [][]int) bool {
	n, p := len(points), 0
	for i := 0; i < n; i++ {
		dx1 := points[(i+1)%n][0] - points[i][0]
		dy1 := points[(i+1)%n][1] - points[i][1]
        dx2 := points[(i+2)%n][0] - points[(i+1)%n][0]
		dy2 := points[(i+2)%n][1] - points[(i+1)%n][1]
        // ����͹����ε����ʣ����бߵ������������������ͬ��
        // ����˽�������Ƿ�仯
        // һ��͹����Σ���֤���ڱ����ǵ�������˳ʱ�����ʱ����ת����������������ж��Ƿ񵥵���ת��
		c := dx1*dy2 - dx2*dy1
		if c != 0 {
			if c*p < 0 {
				return false
			}
			p = c
		}
	}
	return true
}