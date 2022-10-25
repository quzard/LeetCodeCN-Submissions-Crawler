import "math"
func maximumDetonation1(bombs [][]int) int {
    var length  func(x1, y1, x2, y2 int) float64
    length = func(x1, y1, x2, y2 int) float64{
        return math.Sqrt(float64((x2 - x1)*(x2 - x1) + (y2 - y1)*(y2 - y1)))
    }
    
    hashTable := map[int][]int{}
    res := 1
    
    
    for i := 0; i < len(bombs); i++{
        for j := i + 1; j < len(bombs); j++{
            d := length(bombs[i][0], bombs[i][1], bombs[j][0], bombs[j][1])
            if float64(bombs[i][2]) >= d{
                hashTable[i] = append(hashTable[i], j)
            }
            if float64(bombs[j][2]) >= d{
                hashTable[j] = append(hashTable[j], i)
            }
        }
    }
    
    
    for key, _ := range hashTable{
        selected := map[int]bool{}
        if selected[key] == false{
            selected[key] = true
            list := hashTable[key]
            nums := 1
            for len(list) != 0{
                newList := []int{}
                for i := 0; i < len(list); i++{
                    if selected[list[i]] == false{
                        selected[list[i]] = true
                        nums++
                        if len(hashTable[list[i]]) > 0{
                            newList = append(newList, hashTable[list[i]]...)
                        }
                    }
                }
                list = newList
            }
            if res < nums{
                res = nums
            }
        }
    }
    return res
}


func maximumDetonation(a [][]int) (ans int) {
	n := len(a)
	g := make([][]int, n)

	for i, v := range a {
		for j, w := range a {
			if j == i {
				continue
			}
			if (w[0] - v[0]) * (w[0] - v[0]) + (w[1]-v[1])*(w[1]-v[1]) <= v[2]*v[2] {
				g[i] = append(g[i], j)
			}
		}
	}

	for i := range g {
		vis := make([]bool, len(g))
		c:=0
		var dfs func(i int)
		dfs = func(i int) {
			vis[i] = true
			c++
			for _, w := range g[i] {
				if !vis[w] {
					dfs(w)
				}
			}
			return
		}
		dfs(i)
		if c > ans {
			ans =c
		}
	}
	return
}

 