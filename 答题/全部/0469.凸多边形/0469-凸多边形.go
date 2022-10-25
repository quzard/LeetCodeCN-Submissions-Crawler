func isConvex(points [][]int) bool {
	n, p := len(points), 0
	for i := 0; i < n; i++ {
		dx1 := points[(i+1)%n][0] - points[i][0]
		dy1 := points[(i+1)%n][1] - points[i][1]
        dx2 := points[(i+2)%n][0] - points[(i+1)%n][0]
		dy2 := points[(i+2)%n][1] - points[(i+1)%n][1]
        // 利用凸多边形的性质，所有边的向量叉积（向量积）同号
        // 看叉乘结果方向是否变化
        // 一个凸多边形，保证相邻边总是单调的向顺时针或逆时针旋转。叉积的正负可以判断是否单调旋转。
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