//���鼯
func countComponents(n int, edges [][]int) int {
	ans := 0
	fa := make([]int, n)
	//MakeSet �������鼯
	for i := 0; i < n; i++ {
		fa[i] = i
	}

	for _, edge := range edges {
		x := edge[0]
		y := edge[1]

		unionSet(fa, x, y)
	}

	for i := 0; i < n; i++ {
		//�������鼯�ĸ����𰸼�1
		if Find(fa, i) == i {
			ans++
		}
	}
    fmt.Println(fa)
	return ans
}

//���鼯 ·��ѹ��ģ��
//��ѯ��Find������ѯ����Ԫ���Ƿ���ͬһ�������С�
func Find(fa []int, x int) int {
    // һ��ʼx ָ��x
	if fa[x] == x {
		return x
	}
    // �Ҹ��ڵ�
	fa[x] = Find(fa, fa[x])
	return Find(fa, fa[x])
}

func unionSet(fa []int, x, y int) {
	x, y = Find(fa, x), Find(fa, y)
	if x != y {
        // y ָ�� x, x Ϊ���ڵ㣬 β��ָ��ͷ
		fa[y] = x
	}
}

// ��ͨͼ
func countComponents1(n int, edges [][]int) (cnt int) {
	table := make(map[int][]int, n)
	for _, edge := range edges {
		from := edge[0]                       // ͷ
		to := edge[1]                         // β
		table[from] = append(table[from], to) // ͷ����β��
		table[to] = append(table[to], from)   // β�Ͱ���ͷ
	}

	visted := make([]bool, n)

	var dfs func(int)
	dfs = func(index int) {
		visted[index] = true
		for _, node := range table[index] {
			if !visted[node] {
				dfs(node)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visted[i] {
			cnt++
			dfs(i)
		}
	}

	return
}