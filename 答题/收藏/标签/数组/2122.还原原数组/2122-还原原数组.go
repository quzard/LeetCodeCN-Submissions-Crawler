func recoverArray1(nums []int) []int {
    sort.Ints(nums)
    hashTable := map[int]int{}
    hashNums := map[int]bool{}
    for i := 0; i < len(nums); i++{
        hashNums[nums[i]] = true
        for j := i + 1; j < len(nums); j++{
            hashTable[abs(nums[i], nums[j])]++
        }
    }
    res := []int{}
    for key, value := range hashTable{
        if value < len(nums) / 2 || key == 0{
            continue
        }
        res = []int{}
        visited := map[int]bool{}
        for i := 0; i < len(nums); i++{
            if visited[i]{
                continue
            }
            visited[i] = true
            for j := i + 1; j < len(nums); j++{
                if nums[j] == nums[i] + key * 2 && visited[j] == false{
                    visited[j] = true
                    res = append(res, nums[i] + key)
                    if len(res) == len(nums) / 2{
                        return res
                    }
                    break
                }
            }
        }
        
    }
    return res
}

func abs(a, b int) int{
    if a > b{
        if (a - b) % 2 > 0{
            return 0
        }
        return (a - b) / 2
    }
    if (b - a) % 2 > 0{
        return 0
    }
    return (b - a) / 2
}

// github.com/EndlessCheng/codeforces-go
func recoverArray(a []int) (ans []int) {
    sort.Ints(a)
    n := len(a)
    if n == 2 {
        return []int{(a[0] + a[1]) / 2}
    }
o:
    for i := 1; i < n; i++ {
        se := a[i]
        if se == a[0] || (se-a[0])&1 > 0 {
            continue
        }

        k := (se - a[0]) / 2

        i, j := 0, i
        vis := make([]bool, n)
        vis[i] = true
        vis[j] = true
        aa := []int{(a[i] + a[j]) / 2}
        for i < n && j < n {
            for i < n && vis[i] {
                i++
            }
            for j < n {
                if vis[j] {
                    j++
                    continue
                }
                if (a[j]-a[i])/2 > k {
                    continue o
                }
                if (a[j]-a[i])&1 == 0 && (a[j]-a[i])/2 == k {
                    break
                }
                j++
            }
            if i == n || j == n {
                break
            }
            vis[i] = true
            vis[j] = true
            aa = append(aa, (a[i]+a[j])/2)
        }
        if len(aa) == n/2 {
            ans = aa
            return
        }
    }
    return
}
