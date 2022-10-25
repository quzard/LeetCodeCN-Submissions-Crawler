// �� 1 �� N���ĳ� 0 �� N - 1��������������
func minimumSemesters(N int, relations [][]int) int {
    preClassCount := map[int]int{}
    nextClasses := make([][]int, N)
    for i := 0; i < N; i++ {
        preClassCount[i] = 0
    }
    for _, r := range relations {
        nextClasses[r[0]-1] = append(nextClasses[r[0]-1], r[1]-1) // ��¼���ÿγ�
        preClassCount[r[1]-1]++ // �����ʼ���
    }
    term := 0
    for len(preClassCount) > 0 {
        term++
        learn := []int{}
        for class, count := range preClassCount {
            if count == 0 {
                learn = append(learn, class) // ���Ϊ 0������ѧϰ
            }
        }
        if len(learn) == 0 { // û�пγ̿���ѧ�ˣ�˵����ѭ��
            return -1
        }
        for _, l := range learn {
            delete(preClassCount, l)
            for _, class := range nextClasses[l] {
                preClassCount[class]--
            }
        }
    }
    return term
}

func minimumSemesters1(n int, relations [][]int) int {
	fa := make([]int, n)
	son := make([][]int, n)
	for _, relation := range relations {
		u, v := relation[0]-1, relation[1]-1
		fa[v]++
		son[u] = append(son[u], v)
	}
	q := make([]int, 0)
	for i := 0; i < n; i++ {
		if fa[i] == 0 {
			q = append(q, i)
		}
	}
	count := 0
	for len(q) > 0 {
		count++
		q2 := make([]int, 0)
		for len(q) > 0 {
			now := q[0]
			q = q[1:]
			for _, v := range son[now] {
				fa[v]--
				if fa[v] == 0 {
					q2 = append(q2, v)
				}
			}
		}
		q = q2
	}
	for i := 0; i < n; i++ {
		if fa[i] != 0 {
			return -1
		}
	}
	return count
}