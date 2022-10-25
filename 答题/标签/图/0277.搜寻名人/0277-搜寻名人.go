/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution1(knows func(a int, b int) bool) func(n int) int {
    return func(n int) int {
        famous := []int{}
        for i := 1; i < n; i++{
            if knows(0, i){
                famous = append(famous, i)
            }
        }
        if len(famous) == 0{
            for i := 1; i < n; i++{
                if ! knows(i, 0){
                    return -1
                }
            }
            return 0
        }else{
            for i := 0; i < len(famous); i++{
                cnt := 0
                for j := 0; j < n; j++{
                    if famous[i] == j{
                        continue
                    }
                    if knows(famous[i], j) || knows(j, famous[i]) == false{
                        break
                    }
                    cnt++
                }
                if cnt == n - 1{
                    return famous[i]
                }
            }
            return -1
        }
        return -1
    }
}

func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		result := 0
		for i := 1; i < n; i++ {
			if knows(result, i) {
				result = i
			}
		}
		for i := 0; i < n; i++ {
			if i != result && (!knows(i, result) || knows(result, i)) {
				return -1
			}
		}
		return result
	}
}