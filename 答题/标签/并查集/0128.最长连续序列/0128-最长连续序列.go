// 并查集
func longestConsecutive1(nums []int) int {
    n := len(nums)
    if n < 2 {
        return n
    }
    fa := make([][]int, n)
    for i := 0; i < n; i++ {
        fa[i] = []int{i, 1}
    }

    ans := 1
    hashTable := map[int]int{}
    for i := 0; i < n; i++ {
        if _, ok := hashTable[nums[i]]; ok {
            continue
        }
        
        num := nums[i]
        if left, ok := hashTable[num-1]; ok {
            ans = max(Union(fa, left, i), ans)
        }
        if right, ok := hashTable[num+1]; ok {
            ans = max(Union(fa, i, right), ans)
        }
        hashTable[num] = i
    }
    return ans
}

func Find(fa [][]int, i int) int {
    if fa[i][0] == i {
        return i
    }
    root := Find(fa, fa[i][0])
    fa[i][0] = root
    return root
}

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


func longestConsecutive(nums []int) int {
    m := make(map[int]bool)
    for _, v := range nums {
        m[v] = true
    }

    var max = 0
    for k, _ := range m {
        // 过滤重复计算, 只计算最左端的数
        if m[k-1] {
            continue
        }
        cur := k
        step := 1
        for true {
            if m[cur] == false {
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