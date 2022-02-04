func longestConsecutive1(nums []int) int {
	res := 0
	hashTable := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {
		hashTable[nums[i]] = struct{}{}
	}
	for key, _ := range hashTable {
		left := key - 1
		right := key + 1
		length := 1
		delete(hashTable, key)
		for {
			isExited := false
			if _, ok := hashTable[left]; ok {
				delete(hashTable, left)
				length++
				left--
				isExited = true
			}
			if _, ok := hashTable[right]; ok {
				delete(hashTable, right)
				length++
				right++
				isExited = true
			}
			if isExited == false {
				break
			}
		}
		if res < length {
			res = length
		}
	}
	return res
}

func longestConsecutive2(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v] = 0
	}

	var max = 0
	for k, _ := range m {
		if _, ok := m[k-1]; ok { // 过滤重复计算
			continue
		}
		cur := k
		step := 1
		for true {
			if _, ok := m[cur]; !ok {
				break
			}

			if step > max {
				max = step
			}

			cur++
			step++
		}
	}

	return max
}


func Find(fa [][]int, i int) int{
    if fa[i][0] == i {
        return i
    }
    root := Find(fa, fa[i][0])
    fa[i][0] = root
    return root
}

// nums[i] == nums[j] - 1
func Union(fa [][]int, i, j int) int {
    x, y := Find(fa, i), Find(fa, j)
    if x != y {
        fa[x][0] = y
        fa[y][1] += fa[x][1]
    }
    return fa[y][1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// 并查集
func longestConsecutive(nums []int) int {
    n := len(nums)
    if n < 2 {
        return n
    }
    fa := make([][]int, n)
    ans := 1
    hashTable := map[int]int{}
    for i := 0; i < n; i++ {
        fa[i] = []int{i, 1}
    }
    for i := 0; i < n; i++ {
        if _, ok := hashTable[nums[i]]; ok {
            continue
        }
        num := nums[i]
        if left, ok := hashTable[num - 1]; ok {
            ans = max(Union(fa, left, i), ans)
        }
        if right, ok := hashTable[num + 1]; ok {
            ans = max(Union(fa, i, right), ans)
        }
        hashTable[num] = i
    }
    return ans
}