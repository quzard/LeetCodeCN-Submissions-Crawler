//并查集
func countComponents(n int, edges [][]int) int {
	ans := 0
	fa := make([]int, n)
	//MakeSet 建立并查集
	for i := 0; i < n; i++ {
		fa[i] = i
	}

	for _, edge := range edges {
		x := edge[0]
		y := edge[1]

		unionSet(fa, x, y)
	}

	for i := 0; i < n; i++ {
		//遇到并查集的根，答案加1
		if Find(fa, i) == i {
			ans++
		}
	}
    fmt.Println(fa)
	return ans
}

//并查集 路径压缩模板
//查询（Find）：查询两个元素是否在同一个集合中。
func Find(fa []int, x int) int {
    // 一开始x 指向x
	if fa[x] == x {
		return x
	}
    // 找根节点
	fa[x] = Find(fa, fa[x])
	return Find(fa, fa[x])
}

func unionSet(fa []int, x, y int) {
	x, y = Find(fa, x), Find(fa, y)
	if x != y {
        // y 指向 x, x 为父节点， 尾巴指向头
		fa[y] = x
	}
}

// 连通图
func countComponents1(n int, edges [][]int) (cnt int) {
	table := make(map[int][]int, n)
	for _, edge := range edges {
		from := edge[0]                       // 头
		to := edge[1]                         // 尾
		table[from] = append(table[from], to) // 头包含尾巴
		table[to] = append(table[to], from)   // 尾巴包含头
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