func findAllRecipes1(recipes []string, ingredients [][]string, supplies []string) []string {
    res := []string{}
    hashTable := map[string]int{}
    for index, recipe := range recipes{
        hashTable[recipe] = index + 1
    }

    selected := map[string]bool{}
    for _, supplie := range supplies{
        selected[supplie] = true
    }
    
    s := map[string]bool{}
    for i := 0; i < len(recipes); i++{
        if s[recipes[i]] {
            continue
        }
        visited := map[string]bool{}
        visited[recipes[i]] = true
        ingredient := ingredients[i]
        r := []string{recipes[i]}
        find := true
        for len(ingredient) > 0{
            value := ingredient[0]
            if hashTable[value] > 0{
                if s[value] == false{
                    if visited[value]{
                        find = false
                        break
                    }
                    visited[value] = true
                    r = append(r, value)
                    ingredient = append(ingredient, ingredients[hashTable[value] - 1]...)
                }
            }else if selected[value] == false{
                find = false
                break
            }
            ingredient = ingredient[1:]
        }
        if find{
            for _, value := range r{
                if s[value] == false{
                    res = append(res, value)
                }
                s[value] = true
            }
        }
    }
    return res
}

// 拓扑排序
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) (ans []string) {
	g := map[string][]string{}
	deg := map[string]int{}
	for i, r := range recipes {
		for _, s := range ingredients[i] {
			g[s] = append(g[s], r) // 从这道菜的原材料向这道菜连边
			deg[r]++
		}
	}
	for len(supplies) > 0 { // 拓扑排序
		s := supplies[0]
		supplies = supplies[1:]
		for _, r := range g[s] {
			if deg[r]--; deg[r] == 0 { // 这道菜的所有原材料我们都有
				supplies = append(supplies, r)
				ans = append(ans, r)
			}
		}
	}
	return
}